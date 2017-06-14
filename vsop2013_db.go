// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

import (
	"fmt"
	"sync"
	"time"
	"errors"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plb97/vsop2013/configuration"
)

func getCoefs(planet Planet, jd float64) (*planets_t,*coefs_t, error) {
	db := db_pool.Get().(*sql.DB)
	defer db_pool.Put(db)
	
	var p = planets_t{id:int(planet)}
	rowsp, err := db.Query("SELECT name,t1,t2,step,ncf FROM vsop2013_planets WHERE id = ?",p.id)
	if err != nil {
		return nil,nil, err
	}
	defer rowsp.Close()
	for rowsp.Next() {
		err := rowsp.Scan(&p.name, &p.t1, &p.t2, &p.step, &p.ncf)
		if nil != err {
			return nil, nil, err
		}
	}

	var c = coefs_t{jd0:(int(jd - t0 + 0.1) / p.step) * p.step, planet_id:int(planet),cfs:make([][6]float64,p.ncf)}
//fmt.Printf("getCoefs: jd0=%f dt=%d planet_id=%d\n",t0+float64(c.dt),c.dt,c.planet_id)
	rowsc, err := db.Query("SELECT cf,c0,c1,c2,c3,c4,c5 FROM vsop2013_coefs WHERE jd0 = ? AND planet_id = ?",c.jd0,c.planet_id)
	if err != nil {
		return nil, nil, err
	}
	defer rowsc.Close()
	var (
		cf int
		c0,c1,c2,c3,c4,c5 float64
	)
	for rowsc.Next() {
		err := rowsc.Scan(&cf, &c0, &c1, &c2, &c3, &c4, &c5)
		if nil != err {
			return nil, nil, err
		}
		c.cfs[cf] = [6]float64{c0,c1,c2,c3,c4,c5}
	}
	return &p,&c, nil 
}

func EphVsop2013_db(planet Planet, jd float64) (*Ephemeris_t, error) {

//fmt.Println("EphVsop2013_db")
	p, coefs, err := getCoefs(planet,jd)
	if nil != err {
		return nil, err
	}
	if jd < p.t1 {
		return nil, errors.New(fmt.Sprintf("Date error: jd %f < %f",jd,p.t1))
	}
	if p.t2 < jd {
		return nil, errors.New(fmt.Sprintf("Date error: jd %f > %f",jd,p.t2))
	}
//fmt.Printf("coefs=%v\n",coefs)
	djd := jd - t0
	dt0 := (int(djd + 0.1) / p.step) * p.step
//fmt.Printf("dt0=%d djd=%f\n",dt0,djd)
	var tn = make([]float64,p.ncf)
	x:=2.e0*(djd-float64(dt0))/float64(p.step) - 1.e0
//fmt.Printf("x=%f\n",x)	
	tn[0]=1.e0
	tn[1]=x
	for i := 2; i < p.ncf; i++ {
		tn[i]=2.e0*x*tn[i-1]-tn[i-2]
	}
//fmt.Printf("EphVsop2013_db: tn=%v\n",tn)
//fmt.Printf("EphVsop2013_db: cfs=%v\n",coefs.cfs)
	
	r := &Ephemeris_t{p:planet,jd:jd}
	for i := 0; i < 6; i++ {
		r.r[i]=0.e0
		for j := 0; j < p.ncf; j++ {
			r.r[i]+=tn[j]*coefs.cfs[j][i]
//if 0 == i {
//	fmt.Printf("EphVsop2013_db: tn[%d]=%v coefs[%d][%d]=%v r[%d]=%v\n",j,tn[j],j,i,coefs.cfs[j][i],i,r[i])
//}
		}
	}

	return r,nil
}

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

DROP TABLE IF EXISTS vsop2013_coefs;
CREATE TABLE IF NOT EXISTS vsop2013_coefs (
  jd0 int(11) NOT NULL,
  planet_id int(11) NOT NULL,
  cf int(11) NOT NULL,
  c0 double NOT NULL,
  c1 double NOT NULL,
  c2 double NOT NULL,
  c3 double NOT NULL,
  c4 double NOT NULL,
  c5 double NOT NULL,
  PRIMARY KEY (jd0,planet_id,cf)
) 
ENGINE=InnoDB
DEFAULT CHARSET=latin1
PARTITION BY RANGE (jd0) (
    PARTITION m4000 VALUES LESS THAN (547904),
    PARTITION m2000 VALUES LESS THAN (1095808),
    PARTITION m1000 VALUES LESS THAN (1643712),
    PARTITION p1000 VALUES LESS THAN (2191616),
    PARTITION p2000 VALUES LESS THAN (2739520),
    PARTITION p4000 VALUES LESS THAN (3287424)
)
;
*/

//func Generate_db2() {
//	db := db_pool.Get().(*sql.DB)
//	defer db_pool.Put(db)
//
//	var istmt, dstmt, ustmt, istmt2, dstmt2 *sql.Stmt
//	var err_stmt error
//	
//	dstmt, err_stmt = db.Prepare("DELETE FROM vsop2013_planets WHERE id = ?")
//	if err_stmt != nil {
//		log.Fatal(err_stmt)
//	}
//	defer dstmt.Close()
//	istmt, err_stmt = db.Prepare("INSERT INTO vsop2013_planets(id, name, t1, t2, step, ncf) VALUES(?,?,?,?,?,?)")
//	if err_stmt != nil {
//		log.Fatal(err_stmt)
//	}
//	defer istmt.Close()
//	ustmt, err_stmt = db.Prepare("UPDATE vsop2013_planets SET t2 = ? WHERE id = ? AND t2 < ?")
//	if err_stmt != nil {
//		log.Fatal(err_stmt)
//	}
//	defer ustmt.Close()
//
//	dstmt2, err_stmt = db.Prepare("DELETE FROM vsop2013_coefs WHERE jd0 = ? AND planet_id = ? AND cf = ?")
//	if err_stmt != nil {
//		log.Fatal(err_stmt)
//	}
//	defer dstmt2.Close()
//	istmt2, err_stmt = db.Prepare("INSERT INTO vsop2013_coefs(jd0, planet_id, cf, c0, c1, c2, c3, c4, c5) VALUES (?,?,?,?,?,?,?,?,?)")
//	if err_stmt != nil {
//		log.Fatal(err_stmt)
//	}
//	defer istmt2.Close()
//	
//	param_size := Sizeof_vsop2013_param_t()
//	bparam := make([]byte,param_size)
//	coef_size := Sizeof_vsop2013_coef_t()
//	bcoef := make([]byte,coef_size)
////	var writers [9]*os.File
//	
//	minfile := 0
//	maxfile := len(names1)
////	minfile := 4
////	maxfile := 5
//
////fmt.Printf("var vsop2013_coords_data = map[Planet]vsop2013_coords_data_t{\n")
//	var planets [9]*planets_t
//	for ifile := minfile; ifile < maxfile; ifile++ {
//		namef := names1[ifile]+".bin"
//fmt.Println("namef",namef)
//		reader, errr := vsop2013_open(fmt.Sprintf("%s/%s",dir2,namef))
//		if nil != errr {
//			panic(errr)
//		}
//		defer reader.Close()
//		
//		// Reading the Header
//		l1, err1 := reader.ReadAt(bparam,0)
//		if param_size != l1 {
//			panic("File error")
//		}
//		if nil != err1 {
//			panic(err1)
//		}
//		// header file
//		param := New_vsop2013_param_t(bparam)
////fmt.Printf("param=%v\n",param)
//
//		minpla := 0
//		maxpla := 9
////		minpla := 0
////		maxpla := 1
//		if 4 != ifile && 8 < maxpla {
//			maxpla=8
//		}
//		for irec := 0; irec < int(param.nintv); irec++ {
//			off := int64(param_size)+int64(irec)*int64(coef_size)
//			l2, err2 := reader.ReadAt(bcoef,off)
//			if coef_size != l2 {
//				panic("File error")
//			}
//			if nil != err2 {
//				panic(err2)
//			}
//			coefs := New_vsop2013_coef_t(bcoef)
////fmt.Printf("coefs=%v\n",coefs)
////panic("fini")
//			for ipla := minpla ; ipla < maxpla; ipla++ {
//				pla := Planet(ipla+1)
//				iad := int(param.loc.iad[ipla])
//				ncf := int(param.loc.ncf[ipla])
//				nsi := int(param.loc.nsi[ipla])
//				step:=int(param.delta/float64(nsi))
//				nmin := iad - 1
////				nmax := nmin + 6*nsi*ncf
////fmt.Printf("iad=%d ncf=%d nsi=%d nmin=%d nmax=%d\n",iad,ncf,nsi,nmin,nmax)
//				if nil == planets[ipla] {
//					p := planets_t{id:int(pla),name:pla.String(),t1:param.t1,t2:param.t2,step:step,ncf:int(ncf)}
//					planets[ipla] = &p
//					if _, err := dstmt.Exec(p.id); nil != err {
//						log.Fatal(err)
//					}
//					if _, err := istmt.Exec(p.id,p.name,p.t1,p.t2,p.step,p.ncf); nil != err {
//						log.Fatal(err)
//					}
//				}
//				p := planets[ipla]
//				var tab [][6]float64
//				for si := 0; si < nsi; si++ {
//					tab = make([][6]float64,ncf)
//					for co := 0; co < 6; co++ {
//						n := nmin+ncf*(6*si+co)
////fmt.Printf("n=%d cfs=%v\n",n,coefs.coef[n:n+ncf])
//						for cf := 0; cf < ncf; cf++ {
//							tab[cf][co] = coefs.coef[n+cf]
//						}
//					}
//					dt := int(coefs.t1 - t0 + 0.1) + si*step
////fmt.Printf("dt=%d jd0=%.1f\n",dt,t0+float64(dt))
//					for cf := 0; cf < ncf; cf++ {
////fmt.Printf("%v	%v	%v	%v	%v	%v\n",tab[cf][0],tab[cf][1],tab[cf][2],tab[cf][3],tab[cf][4],tab[cf][5])
//						if _, err := dstmt2.Exec(dt, p.id, cf); nil != err {
//							log.Fatal(err)
//						}
////						n := nmin + 6*(si*ncf + cf)
////						n := nmin + 6*(cf*nsi + si)
////						n := nmin + cf
////						cx,cy,cz,cu,cv,cw := coefs.coef[n+0*ncf],coefs.coef[n+1*ncf],coefs.coef[n+2*ncf],coefs.coef[n+3*ncf],coefs.coef[n+4*ncf],coefs.coef[n+5*ncf]
//						if _, err := istmt2.Exec(dt, p.id, cf,tab[cf][0],tab[cf][1],tab[cf][2],tab[cf][3],tab[cf][4],tab[cf][5]); nil != err {
//							log.Fatal(err)
//						}
//					}
////panic("fini")
//				}
//				if int(param.nintv) == irec+1 {
//					p.t2 = param.t2
//					if _, err := ustmt.Exec(p.t2,p.id,p.t2); nil != err {
//						log.Fatal(err)
//					}
//				}
////				var s string
////				fmt.Scanf("%s",&s)
////				fmt.Println(s)
//			}
////break
//		}
////break
//	}
////	fmt.Printf("}\n")
//	return
//}

func gen_planets_db() {
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
			log.Fatal(err)
		}
		defer reader.Close()
		
		param, err := vsop2013_read_param(reader)
		if nil != err {
			log.Fatal(err)
		}
		for ipla := minpla; ipla < maxpla; ipla++ {
			pla := Planet(ipla+1)
			ncf := int(param.loc.ncf[ipla])
			nsi := int(param.loc.nsi[ipla])
			step := param.delta / float64(nsi)
			if 8 == ipla && 4 == ifile || 8 != ipla && 0 == ifile {
				if _, err := dstmtp.Exec(pla); nil != err {
					log.Fatal(err)
				}
				if _, err := istmtp.Exec(pla,pla.String(),param.t1,param.t2,step,ncf); nil != err {
					log.Fatal(err)
				}
			} else if 8 != ipla {
				if _, err := ustmtp.Exec(param.t2,pla,param.t2); nil != err {
					log.Fatal(err)
				}
			}
		}
	}
}

func gen_coefs_db(ifile int, wg *sync.WaitGroup) {

	tstart := time.Now()
fmt.Println(names1[ifile],"...")

	db := db_pool.Get().(*sql.DB)
	defer db_pool.Put(db)

	var err error
	var dstmts, istmts *sql.Stmt
	var ustmts [6]*sql.Stmt

	dstmts,err = db.Prepare("DELETE FROM vsop2013_coefs WHERE jd0 = ? AND planet_id = ?")
	if nil != err {
		log.Fatal(err)
	}
	istmts,err = db.Prepare("INSERT INTO vsop2013_coefs (jd0, planet_id, cf) VALUES (?,?,?)")
	if nil != err {
		log.Fatal(err)
	}
	for co := 0; co < 6; co++ {
		ustmts[co],err = db.Prepare(fmt.Sprintf("UPDATE vsop2013_coefs SET c%d = ? WHERE jd0 = ? AND planet_id = ? AND cf = ?",co))
		if nil != err {
			log.Fatal(err)
		}
	}

	minpla := 0
	maxpla := 9
	namef1 := fmt.Sprintf("%s/%s",configuration.Configuration.InputDir,names1[ifile])
	reader, err := vsop2013_open(namef1)
	if nil != err {
		log.Fatal(err)
	}
	defer reader.Close()

	param, err := vsop2013_read_param(reader)
	if nil != err {
		log.Fatal(err)
	}

	for intv := 0; intv < int(param.nintv); intv++ {
		bdt, _, err := vsop2013_read_intv_dates(reader)
		if nil != err {
			log.Fatal(err)
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
					if _, err := dstmts.Exec(jd0,ipla+1); nil != err {
						log.Fatal(err)
					}
					for cf := 0; cf < ncf; cf++ {
						if _, err := istmts.Exec(jd0,ipla+1,cf); nil != err {
							log.Fatal(err)
						}
					}
				}
				for co := 0; co < 6; co++ {
					for cf := 0; cf < ncf; cf++ {
						c, err := vsop2013_read_coef(reader) 
						if nil != err {
							log.Fatal(err)
						}
						if 8 != ipla || 4 == ifile {
							if _, err := ustmts[co].Exec(c,jd0,ipla+1,cf); nil != err {
								log.Fatal(err)
							}
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

func Generate_db() {
//	gen_planets_db()

	var wg sync.WaitGroup
	for ifile, _ := range names1 {
		wg.Add(1)
		go gen_coefs_db(ifile, &wg)
	}
	wg.Wait()
}
