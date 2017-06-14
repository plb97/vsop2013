// Copyright (c) 2017 plb97.
// All rights reserved.
// Use of this source code is governed by a CeCILL-B_V1
// (BSD-style) license that can be found in the
// LICENCE (French) or LICENSE (English) file.
package vsop2013

// ftp://ftp.imcce.fr/pub/

import (
	"github.com/plb97/vsop2013/configuration"
	"fmt"
	"math"
	"bufio"
	"os"
	"errors"
)

func Vsop2013_binfile() {
	for ifile, _ := range names1 {
		namef1 := fmt.Sprintf("%s/vsop2013/%s",configuration.Configuration.InputDir,names1[ifile])
		fmt.Println(namef1)
		reader, errr := vsop2013_open(namef1)
		if nil != errr {
			panic(errr.Error())
		}
		defer reader.Close()
		nul1 := bufio.NewScanner(reader)
		
		namef2 := fmt.Sprintf("%s/%s.bin",configuration.Configuration.OutputDir,names1[ifile])
		writer, errw := vsop2013_create(namef2)
		if nil != errw {
			panic(errw.Error())
		}
		defer writer.Close()
		
		data := new(vsop2013_data_t)
		data.param.idf,data.param.t1,data.param.t2,data.param.delta,data.param.nintv,data.param.ncoef = vsop2013_read1001(nul1)
		data.param.loc = vsop2013_read1002(nul1)
		writer.Write(data.param.Bytes())
		data.coefs = make([]vsop2013_coef_t,data.param.nintv)
		for k := int32(0); k < data.param.nintv; k++ {
			data.coefs[k].t1, data.coefs[k].t2 = vsop2013_read1003(nul1)
			for n := 0; n < 163; n++ { // 163=978/6
				xcoef, icoef := vsop2013_read1004(nul1)
				for l := 0; l < 6; l++ {
					m:=n*6+l
					data.coefs[k].coef[m]=xcoef[l]*math.Pow10(icoef[l])
				}
			}
			writer.Write(data.coefs[k].Bytes())
		}
	}
	return
}

func Generate_ref() {
	for ifile, _ := range names1 {
		namef1 := fmt.Sprintf("%s/%s",configuration.Configuration.InputDir,names1[ifile])
		fmt.Println(namef1)
		reader, errr := vsop2013_open(namef1)
		if nil != errr {
			panic(errr.Error())
		}
		defer reader.Close()
		nul1 := bufio.NewScanner(reader)
		
		namef2 := fmt.Sprintf("%s/%s.txt",configuration.Configuration.OutputDir,names1[ifile])
		writer, errw := vsop2013_create(namef2)
		if nil != errw {
			panic(errw.Error())
		}
		defer writer.Close()
		
		var param vsop2013_param_t
		param.idf,param.t1,param.t2,param.delta,param.nintv,param.ncoef = vsop2013_read1001(nul1)
		param.loc = vsop2013_read1002(nul1)
		fmt.Fprintf(writer,"param=%v\n",param)
		var coef vsop2013_coef_t
		for k := int32(0); k < param.nintv; k++ {
			coef.t1, coef.t2 = vsop2013_read1003(nul1)
			for n := 0; n < 163; n++ {
				xcoef, icoef := vsop2013_read1004(nul1)
				for l := 0; l < 6; l++ {
					m:=n*6+l
					coef.coef[m]=xcoef[l]*math.Pow10(icoef[l])
				}
			}
			fmt.Fprintf(writer,"coef%d=%v\n",k,coef)
			if 2 == k {break}
		}
		break
	}
	return
}

func Compute_ref() {
	var tzeros = [6]float64{
          77432.5e0,     // -4500/01/01 0h
          625307.5e0,    // -3000/01/01 0h
          1173182.5e0,   // -1500/01/01 0h
          1721057.5e0,   //  0000/01/01 0h
          2268932.5e0,   //  1500/01/01 0h
          2816818.5e0,   //  3000/02/01 0h
	}
var periods = []string{	"*** -4500 -3000",
							"*** -3000 -1500",
							"*** -1500 0000",
							"*** 0000 1500",
							"*** 1500 3000",
							"*** 3000 4500",
	} 
	out, erro := vsop2013_create(fmt.Sprintf("%s/VSOP2013_compute.txt",configuration.Configuration.OutputDir))
	if nil != erro {
		panic(erro.Error())
	}
	defer out.Close()

	ndat:=5
	step:=136798.e0

	for ifile := 0; ifile < len(names1); ifile++ {
		namef1 := fmt.Sprintf("%s.bin",names1[ifile])
		reader, errr := vsop2013_open(fmt.Sprintf("%s/%s",configuration.Configuration.OutputDir,namef1))
		if nil != errr {
			panic(errr.Error())
		}
		defer reader.Close()
			
		writer, errw := vsop2013_create(fmt.Sprintf("%s/%s.compute.txt",configuration.Configuration.OutputDir,names1[ifile]))
		if nil != errw {
			panic(errw.Error())
		}
		defer writer.Close()
		
		tzero := tzeros[ifile]

//fmt.Println("ifile",ifile,"namef1",namef1,"namef2",namef2)
		fmt.Fprintf(out,"\r\n")
		fmt.Fprintf(out,"  %s\r\n",periods[ifile])
		fmt.Fprintf(out,"\r\n")
		maxpla := 9
		if 4 != ifile && 8 < maxpla {
			maxpla=8
		}
		for i := 0; i < maxpla; i++ {
			planet := Planet(i+1)
			for n := 0; n < ndat; n++ {
				jd := tzero+float64(n)*step
				t := jd-jd2000
				r,err := ephVsop2013_ref (t,planet,reader)
				if nil != err {
					panic(err.Error())
				}
				fmt.Fprintf(out,"  %-10s  JD%9.1f  X :%16.12f ua    Y :%16.12f ua    Z :%16.12f ua  \r\n"+
					        "                           X':%16.12f ua/d  Y':%16.12f ua/d  Z':%16.12f ua/d\r\n",
					        planet,jd,r[0],r[1],r[2],r[3],r[4],r[5])
				fmt.Fprintf(writer,"%-10s  JD%9.1f  X :%16.12f ua    Y :%16.12f ua    Z :%16.12f ua  \n"+
					        "                         X':%16.12f ua/d  Y':%16.12f ua/d  Z':%16.12f ua/d\n",
					        planet,jd,r[0],r[1],r[2],r[3],r[4],r[5])
			}
		}
	}
}

func ephVsop2013_ref (djb float64, planet Planet, reader *os.File) (*[6]float64, error) {
	var (
		err_date=errors.New("Error Date")
		err_file=errors.New("Error File")
	//	err_planet=errors.New("Error Planet")
	)
	var r = new([6]float64)
//fmt.Println("ephVsop2013_ref")	
//fmt.Printf("djb=%f\n",djb)

	param_size := Sizeof_vsop2013_param_t()
	bparam := make([]byte,param_size)
	coef_size := Sizeof_vsop2013_coef_t()
	bcoef := make([]byte,coef_size)

	// Reading the Header
	l1, err1 := reader.ReadAt(bparam,0)
	if param_size != l1 {
		return nil, err_file
	}
	if nil != err1 {
		return nil, err1
	}
	param := New_vsop2013_param_t(bparam)
	t1 := param.t1-jd2000
	t2 := param.t2-jd2000
//fmt.Printf("param=%v t1=%f t2=%f\n",param,t1,t2)

	// Error in the parameters
	if djb < t1 || djb > t2 {
		return nil, err_date
	}

	iad := param.loc.iad[planet-1]
	ncf := param.loc.ncf[planet-1]
	nsi := param.loc.nsi[planet-1]
	step:= param.delta/float64(nsi)
//fmt.Printf("iad=%d ncf=%d nsi=%d delta2=%f\n",iad,ncf,nsi,delta2)

	// Reading the Chebychev coefficients
	irec:=int64((djb-t1)/param.delta)
	if djb == t2 {
		irec=int64(param.nintv)-1
	}
	off := int64(param_size)+irec*int64(coef_size)
	l2, err2 := reader.ReadAt(bcoef,off)
	if coef_size != l2 {
		return nil, err_file
	}
	if nil != err2 {
		return nil, err2
	}
	coef := New_vsop2013_coef_t(bcoef)
	dj1 := coef.t1-jd2000
	ik:=int32((djb-dj1)/step)
	if ik == nsi {
		ik--
	}
	iloc:=iad+6*ncf*ik
//fmt.Printf("dj1=%f dj2=%f irec=%d coef=%.17f\n",coef.t1,coef.t2,irec,coef.coef[iloc:iloc+6])
	dj0:=dj1+float64(ik)*step
	x:=2.e0*(djb-dj0)/step - 1.e0
//fmt.Printf("ik=%d iloc=%d dj0=%f x=%f\n",ik,iloc,dj0,x)
	var tn = make([]float64,ncf)

	tn[0]=1.e0
	tn[1]=x
	for i := int32(2); i < ncf; i++ {
		tn[i]=2.e0*x*tn[i-1]-tn[i-2]
	}
//fmt.Printf("ephVsop2013_ref: tn=%v\n",tn)
//fmt.Printf("ephVsop2013_ref: coef=%v\n",coef.coef[iloc:iloc+6*ncf])
	
	for i := int32(0); i < 6; i++ {
		r[i]=0.e0
		for j := int32(0); j < ncf; j++ {
			jp:=ncf-j-1
			jt:=iloc+ncf*i+jp-1
			r[i]+=tn[jp]*coef.coef[jt]
//if 0 == i {
//fmt.Printf("ephVsop2013_ref: tn[%d]=%v coef[%d]=%v r[%d]=%v\n",jp,tn[jp],jt,coef.coef[jt],i,r[i])
//}
		}
	}

	return r,nil
}

func EphVsop2013_ref(jds []float64, planets []Planet) (map[float64]map[Planet]*[6]float64, *[]error) {
	var namef string
	var errs *[]error = nil
	var result map[float64]map[Planet]*[6]float64 = make(map[float64]map[Planet]*[6]float64,len(jds))
	for _, jd := range jds {
		switch {
			case 77294.5 <= jd && jd < 625198.5:
				namef = "VSOP2013.m4000.bin"	// [-4500 -3500]
			case 625198.5 <= jd && jd < 1173102.5:
				namef = "VSOP2013.m2000.bin"	// [-3500 -1500]
			case 1173102.5 <= jd && jd < 1721006.5:
				namef = "VSOP2013.m1000.bin"	// [-1500  0]
			case 1721006.5 <= jd && jd < 2268910.5:
				namef = "VSOP2013.p1000.bin"	// [ 0  +1500]
			case 2268910.5 <= jd && jd < 2816814.5:
				namef = "VSOP2013.p2000.bin"	// [+1500 +3000]
			case 2816814.5 <= jd && jd < 3364718.5:
				namef = "VSOP2013.p4000.bin"	// [+3000 +4500]
			default:
				if nil == errs {
					errs = new([]error)
					*errs = make([]error,0,len(jds)*len(planets))
				}
				*errs = append(*errs,errors.New(fmt.Sprintf("Error Date: %f",jd)))
				continue	
		}
		reader, errr := vsop2013_open(fmt.Sprintf("%s/%s",configuration.Configuration.OutputDir,namef))
		if nil != errr {
			if nil == errs {
				errs = new([]error)
				*errs = make([]error,0,len(jds)*len(planets))
			}
			*errs = append(*errs,errr)
			continue	
		}
		defer reader.Close()

		if _, ok := result[jd]; !ok {
			result[jd] = make(map[Planet]*[6]float64,len(planets))
		} 

		djb := jd - jd2000
		for _, planet := range planets {
			if PLUTO == planet && "VSOP2013.p2000.bin" != namef {
				continue
			}
			r, err := ephVsop2013_ref(djb, planet, reader)
			if nil != err {
				if nil == errs {
					errs = new([]error)
					*errs = make([]error,0,len(jds)*len(planets))
				}
				*errs = append(*errs,err)
				continue
			}
			result[jd][planet] = r
		}
	}
	return result, errs
}

func EphVsop2013(planets []Planet, jds []float64) (map[Planet]map[float64]*[6]float64, *[]error) {
	res, err := EphVsop2013_ref(jds, planets)
	result := make(map[Planet]map[float64]*[6]float64,len(planets))
	for jd, u := range res {
		for p, v := range u {
			if _, ok := result[p]; !ok {
				result[p] = make(map[float64]*[6]float64,len(jds))
			}
			result[p][jd] = v
		}
	}
	return result, err
}

func EphVsop2013b(planet Planet, jd float64) (*[6]float64, error) {
	res, errs := EphVsop2013_ref([]float64{jd}, []Planet{planet})
	if nil != errs {
		return nil, (*errs)[0]
	}
	return res[jd][planet], nil
}

type planets_t struct {
	id int
	name string
	t1 float64
	t2 float64
	step int
	ncf int
}
type coefs_t struct {
	jd0 int
	planet_id int
	cfs [][6]float64
}
func (r coefs_t) String() string {
	return fmt.Sprintf("[%d %d %v]",r.jd0,r.planet_id,r.cfs)
}
