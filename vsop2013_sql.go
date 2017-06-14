// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

import (
	"github.com/plb97/vsop2013/configuration"
	"fmt"
	"bytes"
	"strings"
	"encoding/binary"
	"sync"
	"time"
	"errors"
//	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/*
DROP TABLE IF EXISTS vsop2013_planets;
CREATE TABLE vsop2013_planets (
  id int(11) NOT NULL,
  name text NOT NULL,
  t1 double NOT NULL,
  t2 double NOT NULL,
  step int(11) NOT NULL,
  ncf int(11) NOT NULL,
  PRIMARY KEY (id)
) 
ENGINE=InnoDB 
DEFAULT CHARSET=latin1
;

DROP TABLE IF EXISTS vsop2013_mercury;
CREATE TABLE vsop2013_mercury (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_venus;
CREATE TABLE vsop2013_venus (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_emb;
CREATE TABLE vsop2013_emb (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_mars;
CREATE TABLE vsop2013_mars (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_jupiter;
CREATE TABLE vsop2013_jupiter (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_saturn;
CREATE TABLE vsop2013_saturn (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_uranus;
CREATE TABLE vsop2013_uranus (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_neptune;
CREATE TABLE vsop2013_neptune (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
DROP TABLE IF EXISTS vsop2013_pluto;
CREATE TABLE vsop2013_pluto (
 jd0 INT NOT NULL ,
 c0 BLOB NULL DEFAULT NULL ,
 c1 BLOB NULL DEFAULT NULL , 
 c2 BLOB NULL DEFAULT NULL , 
 c3 BLOB NULL DEFAULT NULL , 
 c4 BLOB NULL DEFAULT NULL , 
 c5 BLOB NULL DEFAULT NULL , 
 PRIMARY KEY (jd0)
)
 ENGINE = InnoDB
;
*/


func eph_sql(planet Planet) func (jd float64) (*Ephemeris_t, error) {
	db := db_pool.Get().(*sql.DB)
	defer db_pool.Put(db)

	var p planets_t
	sstmtp, err := db.Prepare("SELECT id,name,t1,t2,step,ncf FROM vsop2013_planets WHERE id = ?")
	if nil != err {
		panic(err)
//		log.Fatal(err)
	}
	defer sstmtp.Close()
	if err := sstmtp.QueryRow(planet).Scan(&p.id,&p.name,&p.t1,&p.t2,&p.step,&p.ncf); nil != err {
		panic(err)
//		log.Fatal(err)
	}
	
	return func (jd float64) (*Ephemeris_t, error) {
		if jd < p.t1 {
			return nil, errors.New(fmt.Sprintf("Date error: jd %f < %f",jd,p.t1))
		}
		if p.t2 < jd {
			return nil, errors.New(fmt.Sprintf("Date error: jd %f > %f",jd,p.t2))
		}
		db := db_pool.Get().(*sql.DB)
		defer db_pool.Put(db)

		sstmt, err := db.Prepare(fmt.Sprintf("SELECT c0, c1, c2, c3, c4, c5 FROM vsop2013_%s WHERE jd0 = ?",strings.ToLower(p.name)))
		if nil != err {
			return nil, err
		}
		defer sstmt.Close()
	
		jd0 := (int(jd - t0) / p.step) * p.step
		var tn = make([]float64,p.ncf)
		x:=2.e0*(jd-t0-float64(jd0))/float64(p.step) - 1.e0
	
		tn[0]=1.e0
		tn[1]=x
		for i := 2; i < p.ncf; i++ {
			tn[i]=2.e0*x*tn[i-1]-tn[i-2]
		}

		l := 8 * p.ncf
		var c = [6][]byte{make([]byte,l),make([]byte,l),make([]byte,l),make([]byte,l),make([]byte,l),make([]byte,l),}
		if err := sstmt.QueryRow(jd0).Scan(&c[0], &c[1], &c[2], &c[3], &c[4], &c[5]); nil != err {
			return nil, err
		}
		
		var f float64
		r := &Ephemeris_t{p:planet,jd:jd}
		for i := 0; i < 6; i++ {
			buf := bytes.NewReader(c[i])
			r.r[i]=0.e0
			for j := 0; j < p.ncf; j++ {
				if err := binary.Read(buf, binary.BigEndian, &f); err != nil {
					return nil, err
				}
				r.r[i]+=tn[j]*f
			}
		}
		return r,nil
	}
}

var EphVsop2013_mercury = eph_sql(MERCURY)
var EphVsop2013_venus = eph_sql(VENUS)
var EphVsop2013_emb = eph_sql(EMB)
var EphVsop2013_mars = eph_sql(MARS)
var EphVsop2013_jupiter = eph_sql(JUPITER)
var EphVsop2013_saturn = eph_sql(SATURN)
var EphVsop2013_uranus = eph_sql(URANUS)
var EphVsop2013_neptune = eph_sql(NEPTUNE)
var EphVsop2013_pluto = eph_sql(PLUTO)

var EphsVsop2013_sql = map[Planet]func (jd float64) (*Ephemeris_t, error) {
	MERCURY:EphVsop2013_mercury,
	VENUS:EphVsop2013_venus,
	EMB:EphVsop2013_emb,
	MARS:EphVsop2013_mars,
	JUPITER:EphVsop2013_jupiter,
	SATURN:EphVsop2013_saturn,
	URANUS:EphVsop2013_uranus,
	NEPTUNE:EphVsop2013_neptune,
	PLUTO:EphVsop2013_pluto,
}

func EphVsop2013_sql(planet Planet, jd float64) (*Ephemeris_t, error) {
	return EphsVsop2013_sql[planet](jd)
}

func gen_planets_sql() {
	db := db_pool.Get().(*sql.DB)
	defer db_pool.Put(db)

	dstmtp, err := db.Prepare("DELETE FROM vsop2013_planets WHERE id = ?")
	if err != nil {
		panic(err)
//		log.Fatal(err)
	}
	defer dstmtp.Close()

	istmtp, err := db.Prepare("INSERT INTO vsop2013_planets(id, name, t1, t2, step, ncf) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err)
//		log.Fatal(err)
	}
	defer istmtp.Close()

	ustmtp, err := db.Prepare("UPDATE vsop2013_planets SET t2 = ? WHERE id = ? AND t2 < ?")
	if err != nil {
		panic(err)
//		log.Fatal(err)
	}
	defer ustmtp.Close()

	minpla := 0
	maxpla := 9
	for ifile := 0; ifile < len(names1); ifile++ {
		namef1 := fmt.Sprintf("%s/%s",configuration.Configuration.InputDir,names1[ifile])
		reader, err := vsop2013_open(namef1)
		if nil != err {
			panic(err)
//			log.Fatal(err)
		}
		defer reader.Close()
		
		param, err := vsop2013_read_param(reader)
		if nil != err {
			panic(err)
//			log.Fatal(err)
		}
		for ipla := minpla; ipla < maxpla; ipla++ {
			pla := Planet(ipla+1)
			ncf := int(param.loc.ncf[ipla])
			nsi := int(param.loc.nsi[ipla])
			step := param.delta / float64(nsi)
			if 8 == ipla && 4 == ifile || 8 != ipla && 0 == ifile {
				if _, err := dstmtp.Exec(pla); nil != err {
					panic(err)
//					log.Fatal(err)
				}
				if _, err := istmtp.Exec(pla,pla.String(),param.t1,param.t2,step,ncf); nil != err {
					panic(err)
//					log.Fatal(err)
				}
			} else if 8 != ipla {
				if _, err := ustmtp.Exec(param.t2,pla,param.t2); nil != err {
					panic(err)
//					log.Fatal(err)
				}
			}
		}
	}
}

func gen_sql(ifile int, wg *sync.WaitGroup) {

	tstart := time.Now()
fmt.Println(names1[ifile],"...")

	db := db_pool.Get().(*sql.DB)
	defer db_pool.Put(db)

	var err error
	var dstmts, istmts = make([]*sql.Stmt,9), make([]*sql.Stmt,9)
	var ustmts = [6][]*sql.Stmt{0:make([]*sql.Stmt,9),
								1:make([]*sql.Stmt,9),
								2:make([]*sql.Stmt,9),
								3:make([]*sql.Stmt,9),
								4:make([]*sql.Stmt,9),
								5:make([]*sql.Stmt,9),
	}
	for ipla := 0; ipla < 9; ipla++ {
		pla := strings.ToLower(Planet(ipla+1).String())
		dstmts[ipla],err = db.Prepare(fmt.Sprintf("DELETE FROM vsop2013_%s WHERE jd0 = ?",pla))
		if nil != err {
			panic(err)
//			log.Fatal(err)
		}
		istmts[ipla],err = db.Prepare(fmt.Sprintf("INSERT INTO vsop2013_%s (jd0) VALUES (?)",pla))
		if nil != err {
			panic(err)
//			log.Fatal(err)
		}
		for co := 0; co < 6; co++ {
			ustmts[co][ipla], err = db.Prepare(fmt.Sprintf("UPDATE vsop2013_%s SET c%d = ? WHERE jd0 = ?",pla,co))
			if nil != err {
			panic(err)
//			log.Fatal(err)
			}
		}
	}
	minpla := 0		
	maxpla := 9
	buf := new(bytes.Buffer)	
	namef1 := fmt.Sprintf("%s/%s",configuration.Configuration.InputDir,names1[ifile])
	reader, err := vsop2013_open(namef1)
	if nil != err {
		panic(err)
	}
	defer reader.Close()

	param, err := vsop2013_read_param(reader)
	if nil != err {
		panic(err)
	}
	for intv := 0; intv < int(param.nintv); intv++ {
		bdt, _, err := vsop2013_read_intv_dates(reader)
		if nil != err {
			panic(err)
		}
		for ipla := minpla; ipla < maxpla; ipla++ {
//			pla := Planet(ipla+1)
//fmt.Println("pla",pla)
//			iad := int(param.loc.iad[ipla])
			ncf := int(param.loc.ncf[ipla])
			nsi := int(param.loc.nsi[ipla])
//			nmin := iad - 1
//			nmax := nmin + 6*ncf*nsi
			step := param.delta / float64(nsi)
			for si := 0; si < nsi; si++ {
				dt1 := bdt + float64(si)*step
//				dt2 := dt1 + step
				jd0 := int(dt1 - t0)
				if 8 != ipla || 4 == ifile {
					if _, err := dstmts[ipla].Exec(jd0); nil != err {
						panic(err)
//						log.Fatal(err)
					}
					if _, err := istmts[ipla].Exec(jd0); nil != err {
						panic(err)
//						log.Fatal(err)
					}
				}
				for co := 0; co < 6; co++ {
					buf.Reset()
					for cf := 0; cf < ncf; cf++ {
						c, err := vsop2013_read_coef(reader) 
						if nil != err {
							panic(err)
//							log.Fatal(err)
						}
						binary.Write(buf, binary.BigEndian, c)
					}
					if 8 != ipla || 4 == ifile {
						if _, err := ustmts[co][ipla].Exec(buf.Bytes(),jd0); nil != err {
							panic(err)
//							log.Fatal(err)
						}
					}
				}
			}
		}
	}
	wg.Done()
	tend := time.Now()
fmt.Println("...",names1[ifile], "duration", tend.Sub(tstart))

}

func Generate_sql() {
//	gen_planets_sql()
	
	var wg sync.WaitGroup
 
	for ifile, _ := range names1 {
		wg.Add(1)
		go gen_sql(ifile, &wg)
	}
	wg.Wait()
}
