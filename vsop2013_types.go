// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

const jd2000=2451545.e0
const t0=77294.5e0

type Planet int
const (
	MERCURY	Planet = iota+1
	VENUS
	EMB
	MARS
	JUPITER
	SATURN
	URANUS
	NEPTUNE
	PLUTO
	
	EARTH_MOON_BARYCENTER = EMB
)
var planetcode = [...]string{"MERCURY","VENUS","EMB","MARS","JUPITER","SATURN","URANUS","NEPTUNE","PLUTO"}
func (ip Planet) String() string {
	return planetcode[ip-1]
}

var PLANETS = []Planet {MERCURY, VENUS, EMB, MARS, JUPITER, SATURN, URANUS, NEPTUNE, PLUTO}

type vsop2013_coef_t struct {
	t1, t2 float64
	coef [978]float64
}
func Sizeof_vsop2013_coef_t() int {
	return 2*8+978*8
}
func (r vsop2013_coef_t) String() string {
	line := fmt.Sprintf("vsop2013_coef_t { t1:%f, t2:%f, coef:[978]float64{",r.t1,r.t2)
	for i, v := range r.coef {
		line += fmt.Sprintf("%d:%.15e,",i,v)
	}
	line += "}}"
	return line
}
func (r vsop2013_coef_t) Bytes() []byte {
	var err error
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, r.t1)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.t2)
	if nil != err {
		panic(err.Error())
	}
	for i := 0; i < 978; i++ {
		binary.Write(buf, binary.BigEndian, r.coef[i])
		if nil != err {
			panic(err.Error())
		}
	}
	return buf.Bytes()
}
func New_vsop2013_coef_t(b []byte) vsop2013_coef_t {
	var err error
	buf := bytes.NewReader(b)
	r := new(vsop2013_coef_t)
	err = binary.Read(buf, binary.BigEndian, &r.t1)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.t2)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < 978; i++ {
		err = binary.Read(buf, binary.BigEndian, &r.coef[i])
		if err != nil {
			panic(err.Error())
		}
	}
	return *r
}

type vsop2013_loc_t struct {
	iad [9]int32
	ncf [9]int32
	nsi [9]int32
}
func (r vsop2013_loc_t) String() string {
	line := ""
	line += "vsop2013_loc_t {"
	line += fmt.Sprintf("iad:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.iad[i])
	}
	line += "},"
	line += fmt.Sprintf("ncf:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.ncf[i])
	}
	line += "},"
	line += fmt.Sprintf("nsi:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.nsi[i])
	}
	line += "},"
	line += "}"
	return line
}

type vsop2013_param_t struct {
	idf int32
	t1,t2,delta float64
	nintv,ncoef int32
	loc vsop2013_loc_t
}
func (r vsop2013_param_t) String() string {
	line := ""
	line += fmt.Sprintf("vsop2013_param_t {idf:%d, t1:%f, t2:%f, delta:%f, nintv:%d, ncoef:%d, ",
		r.idf,r.t1,r.t2,r.delta,r.nintv,r.ncoef)
	line += fmt.Sprintf("loc:%v,",r.loc)
	line += "}"
	return line
}
func Sizeof_vsop2013_param_t() int {
	return 4+3*8+2*4+3*9*4
}
func (r vsop2013_param_t) Bytes() []byte {
	var err error
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, r.idf)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.t1)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.t2)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.delta)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.nintv)
	if nil != err {
		panic(err.Error())
	}
	binary.Write(buf, binary.BigEndian, r.ncoef)
	if nil != err {
		panic(err.Error())
	}
	for i := 0; i < 9; i++ {
		binary.Write(buf, binary.BigEndian, r.loc.iad[i])
		if nil != err {
			panic(err.Error())
		}
	}
	for i := 0; i < 9; i++ {
		binary.Write(buf, binary.BigEndian, r.loc.ncf[i])
		if nil != err {
			panic(err.Error())
		}
	}
	for i := 0; i < 9; i++ {
		binary.Write(buf, binary.BigEndian, r.loc.nsi[i])
		if nil != err {
			panic(err.Error())
		}
	}
	return buf.Bytes()
}
func New_vsop2013_param_t(b []byte) vsop2013_param_t {
	var err error
	buf := bytes.NewReader(b)
	r := new(vsop2013_param_t)
	err = binary.Read(buf, binary.BigEndian, &r.idf)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.t1)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.t2)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.delta)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.nintv)
	if err != nil {
		panic(err.Error())
	}
	err = binary.Read(buf, binary.BigEndian, &r.ncoef)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < 9; i++ {
		err = binary.Read(buf, binary.BigEndian, &r.loc.iad[i])
		if err != nil {
			panic(err.Error())
		}
	}
	for i := 0; i < 9; i++ {
		err = binary.Read(buf, binary.BigEndian, &r.loc.ncf[i])
		if err != nil {
			panic(err.Error())
		}
	}
	for i := 0; i < 9; i++ {
		err = binary.Read(buf, binary.BigEndian, &r.loc.nsi[i])
		if err != nil {
			panic(err.Error())
		}
	}
	return *r
}

type vsop2013_data_t struct {
	param vsop2013_param_t
	coefs []vsop2013_coef_t
}
func (r vsop2013_data_t) String() string {
//fmt.Println("sizeof vsop2013_param_t",r.param.sizeof())
//fmt.Println("sizeof vsop2013_coef_t",r.coefs[0].sizeof())
	line := ""
	line += fmt.Sprintf("vsop2013_data_t {\n	param:vsop2013_param_t {idf:%d, t1:%f, t2:%f, delta:%f,nintv:%d, ncoef:%d,\n	loc:vsop2013_loc_t {\n",
				r.param.idf,r.param.t1,r.param.t2,r.param.delta,r.param.nintv,r.param.ncoef)
	line += fmt.Sprintf("		iad:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.param.loc.iad[i])
	}
	line += "},\n"
	line += fmt.Sprintf("		ncf:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.param.loc.ncf[i])
	}
	line += "},\n"
	line += fmt.Sprintf("		nsi:[9]int32{")
	for i:= 0; i < 9; i++ {
		line += fmt.Sprintf("%d,",r.param.loc.nsi[i])
	}
	line += "},\n"
	line += "	}},\n"
	line += "	coefs:[]vsop2013_coef_t {\n"
	for i, v := range r.coefs {
		line += fmt.Sprintf("		%d:%v,\n",i,v)
	}
	line += "	},\n"
	line += "}\n"
	return line
}

type vsop2013_coord_t struct {
	cx, cy, cz, cu, cv, cw float64
}
func (r vsop2013_coord_t) String() string {
	return fmt.Sprintf("vsop2013_coord_t{cx:%.14e,cy:%.14e,cz:%.14e,cu:%.14e,cv:%.14e,cw:%.14e}",r.cx,r.cy,r.cz,r.cu,r.cv,r.cw)
}

type vsop2013_coords_data_t struct {
	planet Planet
	t1 float64
	t2 float64
	step float64
	ncf int
	c [][]vsop2013_coords_t
}

type vsop2013_coords_t struct {
//	pla Planet
	dj1 float64
	dj2 float64
//	ncf int32
	cfs []vsop2013_coord_t
}
func (r vsop2013_coords_t) String() string {
	line := "vsop2013_coords_t{"

	line += fmt.Sprintf("dj1:%f,dj2:%f:%d,csf:[]vsop2013_coord_t{",r.dj1,r.dj2)
	for i := 0; i < len(r.cfs); i++ {
		line += fmt.Sprintf("%v,",r.cfs[i])
	}
	line += "}"

	line += "}"
	return line
}

type Ephemeris_t struct {
	p Planet
	jd float64
	r [6]float64
}
func (r Ephemeris_t) String() string {
	return fmt.Sprintf("[%s %f %v]",r.p,r.jd,r.r)
}
func (r Ephemeris_t) P() string {
	return r.String()
}
func (r Ephemeris_t) JD() float64 {
	return r.jd
}
func (r Ephemeris_t) E() [6]float64 {
	return r.r
}
//func (r Ephemeris_t) X() float64 {
//	return r.r[0]
//}
//func (r Ephemeris_t) Y() float64 {
//	return r.r[1]
//}
//func (r Ephemeris_t) Z() float64 {
//	return r.r[2]
//}
//func (r Ephemeris_t) Vx() float64 {
//	return r.r[3]
//}
//func (r Ephemeris_t) Vy() float64 {
//	return r.r[4]
//}
//func (r Ephemeris_t) Vz() float64 {
//	return r.r[5]
//}
