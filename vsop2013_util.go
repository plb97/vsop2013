// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

import (
	"os"
	"bufio"
	"fmt"
	"io"
	"errors"
	"math"

)

var names1 = []string{	"VSOP2013.m4000",	// [-4500 -3500]	  77294.5   625198.5
			"VSOP2013.m2000",	// [-3500 -1500]	 625198.5  1173102.5
			"VSOP2013.m1000",	// [-1500  0]		1173102.5  1721006.5
			"VSOP2013.p1000",	// [ 0  +1500]		1721006.5  2268910.5
			"VSOP2013.p2000",	// [+1500 +3000]	2268910.5  2816814.5
			"VSOP2013.p4000",	// [+3000 +4500]	2816814.5  3364718.5
}

func vsop2013_open(file string) (*os.File, error) {
	reader, err := os.Open(file)
	if nil != err {
		//panic(err.Error())
		return nil, err
	}
	return reader, nil
}

func vsop2013_create(file string) (*os.File, error) {
	writer, err := os.Create(file)
	if nil != err {
		//panic(err.Error())
		return nil, err
	}
	return writer, nil
}

func vsop2013_read_line(nulog *bufio.Scanner) string {
	if !nulog.Scan() {
		panic("end of file.")
	}
	line := nulog.Text()
	return line
}

func vsop2013_read_param(r io.Reader) (*vsop2013_param_t, error) {
	var r_err = errors.New("read error")
	var n int
	var err error
	var p vsop2013_param_t
	n,err = fmt.Fscanf(r,"%d\n%f\n%f\n%f\n%d\n%d\n",&p.idf,&p.t1,&p.t2,&p.delta,&p.nintv,&p.ncoef)
	if nil != err {
		return nil, err
	}
	if 6 != n {
		return nil,r_err
	}
	n,err = fmt.Fscanf(r,"%d %d %d %d %d %d %d %d %d",
		&p.loc.iad[0],&p.loc.iad[1],&p.loc.iad[2],&p.loc.iad[3],&p.loc.iad[4],&p.loc.iad[5],&p.loc.iad[6],&p.loc.iad[7],&p.loc.iad[8])
	if nil != err {
		return nil, err
	}
	if 9 != n {
		return nil,r_err
	}
	n,err = fmt.Fscanf(r,"%d %d %d %d %d %d %d %d %d",
		&p.loc.ncf[0],&p.loc.ncf[1],&p.loc.ncf[2],&p.loc.ncf[3],&p.loc.ncf[4],&p.loc.ncf[5],&p.loc.ncf[6],&p.loc.ncf[7],&p.loc.ncf[8])
	if nil != err {
		return nil, err
	}
	if 9 != n {
		return nil,r_err
	}
	n,err = fmt.Fscanf(r,"%d %d %d %d %d %d %d %d %d",
		&p.loc.nsi[0],&p.loc.nsi[1],&p.loc.nsi[2],&p.loc.nsi[3],&p.loc.nsi[4],&p.loc.nsi[5],&p.loc.nsi[6],&p.loc.nsi[7],&p.loc.nsi[8])
	if nil != err {
		return nil, err
	}
	if 9 != n {
		return nil,r_err
	}
	return &p,nil
}
func vsop2013_read_intv_dates(r io.Reader) (float64,float64,error) {
	var r_err = errors.New("read error")
	var err error
	var n int
	var bdt, edt float64
	n,err = fmt.Fscanf(r,"%f %f",&bdt,&edt)
	if nil != err {
		return bdt, edt, err
	}
	if 2 != n {
		return bdt, edt,r_err
	}
	return bdt, edt, nil
}
func vsop2013_read_coef(r io.Reader) (float64,error) {
	var r_err = errors.New("read error")
	var err error
	var n int
	var m, c float64
	var e int
	n,err = fmt.Fscanf(r,"%f %d",&m,&e)
	if nil != err {
		return c, err
	}
	if 2 != n {
		return c, r_err
	}
	c = m*math.Pow10(e)
	return c, nil
}

func vsop2013_read1001(nulog *bufio.Scanner) (int32,float64,float64,float64,int32,int32) {
	var t1,t2,delta float64
	var idf,nintv,ncoef int32
	line := vsop2013_read_line(nulog) + " " + // idf
			vsop2013_read_line(nulog) + " " + // t1
			vsop2013_read_line(nulog) + " " + // t2
			vsop2013_read_line(nulog) + " " + // delta
			vsop2013_read_line(nulog) + " " + // nintv
			vsop2013_read_line(nulog) + " "   // ncoef
	fmt.Sscanf(line,"%d %f %f %f %d %d",&idf,&t1,&t2,&delta,&nintv,&ncoef)
	return idf,t1,t2,delta,nintv,ncoef
}

func vsop2013_read1002(nulog *bufio.Scanner) (vsop2013_loc_t) {
	var loc = new(vsop2013_loc_t)
	var line string
	line = vsop2013_read_line(nulog)
	fmt.Sscanf(line,"%d%d%d%d%d%d%d%d%d",&loc.iad[0],&loc.iad[1],&loc.iad[2],&loc.iad[3],&loc.iad[4],&loc.iad[5],&loc.iad[6],&loc.iad[7],&loc.iad[8])
	line = vsop2013_read_line(nulog)
	fmt.Sscanf(line,"%d%d%d%d%d%d%d%d%d",&loc.ncf[0],&loc.ncf[1],&loc.ncf[2],&loc.ncf[3],&loc.ncf[4],&loc.ncf[5],&loc.ncf[6],&loc.ncf[7],&loc.ncf[8])
	line = vsop2013_read_line(nulog)
	fmt.Sscanf(line,"%d%d%d%d%d%d%d%d%d",&loc.nsi[0],&loc.nsi[1],&loc.nsi[2],&loc.nsi[3],&loc.nsi[4],&loc.nsi[5],&loc.nsi[6],&loc.nsi[7],&loc.nsi[8])
	return *loc
}

func vsop2013_read1003(nulog *bufio.Scanner) (float64,float64) {
	var dj1, dj2 float64
	line := vsop2013_read_line(nulog)
	fmt.Sscanf(line,"%f%f%",&dj1,&dj2)
	return dj1, dj2
}

//		icoef [6]int
//		xcoef [6]float64

func vsop2013_read1004(nulog *bufio.Scanner) (*[6]float64, *[6]int) {
	var xcoef = new([6]float64)
	var icoef = new([6]int)
	line := vsop2013_read_line(nulog)
	fmt.Sscanf(line,"%f %d %f %d %f %d %f %d %f %d %f %d",
		&xcoef[0],&icoef[0],&xcoef[1],&icoef[1],&xcoef[2],&icoef[2],&xcoef[3],&icoef[3],&xcoef[4],&icoef[4],&xcoef[5],&icoef[5])
	return xcoef, icoef
}

const ε = math.Pi*(23.0+26.0/60+21.41135/3600)/180 // 23°26'21.41135"
const φ = -math.Pi*(0.05188/3600)/180 // -0.05188"
var hec2heq = [][]float64{
	0:[]float64{0:math.Cos(φ),1:-math.Sin(φ)*math.Cos(ε),2: math.Sin(φ)*math.Sin(ε),},
	1:[]float64{0:math.Sin(φ),1: math.Cos(φ)*math.Cos(ε),2:-math.Cos(φ)*math.Sin(ε),},
	2:[]float64{0:0          ,1: math.Sin(ε)              ,2: math.Cos(ε),},
}
var heq2hec = [][]float64{
	0:[]float64{0: math.Cos(φ)            ,1: math.Sin(φ)            ,2: 0,},
	1:[]float64{0:-math.Sin(φ)*math.Cos(ε),1: math.Cos(φ)*math.Cos(ε),2: math.Sin(ε),},
	2:[]float64{0: math.Sin(φ)*math.Sin(ε),1:-math.Cos(φ)*math.Sin(ε),2: math.Cos(ε),},
}

func HEC2heq(hec *[6]float64) (*[6]float64) {
	var heq [6]float64
	for i := 0; i < 3; i++ {
		heq[i] = 0 
		heq[i+3] = 0
		for j := 0; j < 3; j++ {
			heq[i]   += hec2heq[i][j]*hec[j]
			heq[i+3] += hec2heq[i][j]*hec[j+3]
		}
	}
	return &heq
}

func HEQ2hec(heq *[6]float64) (*[6]float64) {
	var hec [6]float64
	for i := 0; i < 3; i++ {
		hec[i] = 0 
		hec[i+3] = 0
		for j := 0; j < 3; j++ {
			hec[i] += heq2hec[i][j]*heq[j]
			hec[i+3] += heq2hec[i][j]*heq[j+3]
		}
	}
	return &hec
}
func (r Ephemeris_t) Q() [6]float64 {
	return *HEQ2hec(&r.r)
}
