package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"github.com/plb97/vsop2013"
)


//func Datetime(t time.Time) (year int, month int, day float64) {
//	u:= t.UTC()
//	return u.Year(), int(u.Month()), DayFloat(u.Day(),u.Hour(), u.Minute(), u.Second(), u.Nanosecond())
//}
//
//func DayFloat(d, h, m, s, n int) (day float64) {
//	day = float64(d) + float64(h * 3600 + m * 60 + s) / 86400.0  + float64(n) / 86400000000000.0
//	return day
//}
//
//func DayHrMnScNano(d float64) (day, hr, mn, sd, nano int) {
//	x := d
//	day = int(x)
//	x = (x - float64(day))*24
//	hr = int(x)
//	x = (x - float64(hr))*60
//	mn = int(x)
//	x = (x - float64(mn))*60
//	sd = int(x)
//	x = (x - float64(sd))*1000000000
//	nano = int(x+0.5)
//	return day, hr, mn, sd, nano
//}
//
//func TimeToJulian(t time.Time) float64 {
//	y, m, d := Datetime(t)
//	return julian.CalendarGregorianToJD(y,m,d)
//}
//
//func UTC() *time.Location {
//	UTC, err := time.LoadLocation("UTC")
//	if (nil != err) {
//		panic(err)
//	}
//	return UTC
//}
//
//func TZ() *time.Location {
//	return time.Now().Location()
//}
//
//func JulianToTime(jd float64) time.Time {
////	var y, m int
////	var d, ds, hr, mn, sd, na float64
//	y, m, d := julian.JDToCalendar(jd)
////	ds = (d - math.Floor(d)) * 86400.0
////	sd = math.Floor(ds)
////	na = math.Floor((ds-float64(sd))*1000000000.0)
////	hr = math.Floor(sd / 3600.0)
////	sd = sd - hr * 3600.0
////	mn = math.Floor(sd / 60.0)
////	sd = sd - mn * 60.0
//// 
////	day, hr, mn, sd, nano := DayHrMnScNano(d)
////	return time.Date(y,time.Month(m),day,hr,mn,sd,nano,UTC())
//func compareTime(y, m int, d float64) (ct float64) {
//	ct = float64(10000*y+100*m)+d
//	return ct
//}
//
//// In time package the calendrical calculations always assume a Gregorian calendar.
//func TimeToJD(t time.Time) (jd float64) {
//	u := t.UTC()
//	julian := u.Before(time.Date(1582,time.Month(10),15,0,0,0,0,UTC()))
//	Y := u.Year()
//	M := int(u.Month())
//	D :=  DayFloat(u.Day(),u.Hour(), u.Minute(), u.Second(), u.Nanosecond())
//	if M < 3 {
//		Y--
//		M += 12
//	}
//	A := Y / 100
//	B := 2 - A + A/4
//	if julian {
//		B = 0
//	}
//	// JD = INT(365.25*(Y+4716)) + INT(30.6001*(M+1)) + D + B - 1524.5
//	jd = float64((36525*(Y+4716))/100) + float64((306*(M+1))/10 + B) + D - 1524.5
//	return jd
//}
//
//func JDToTime(jd float64) (t time.Time) {
//	return t
//}
//
//}

//func phases() {
//	fmt.Println("phases")
//	TZ := TZ()
//	année:=2016
//	for mois := 1; mois <= 13; mois++ {
//		yd := float64(année)+(float64(mois)-1.3)/12.0
//		nl := JulianToTime(moonphase.New(yd)).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		if (année != nl.Year()) {
////			nl = JulianToTime(moonphase.New(yd+0.04*float64(année-nl.Year()))).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		}
//		pq:=JulianToTime(moonphase.First(yd)).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		if (année != pq.Year()) {
////			pq = JulianToTime(moonphase.First(yd+0.04*float64(année-pq.Year()))).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		}
//		pl:=JulianToTime(moonphase.Full(yd)).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		if (année != pl.Year()) {
////			pl = JulianToTime(moonphase.Full(yd+0.04*float64(année-pl.Year()))).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		}
//		dq:=JulianToTime(moonphase.Last(yd)).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		if (année != dq.Year()) {
////			dq = JulianToTime(moonphase.Last(yd+0.04*float64(année-dq.Year()))).In(TZ).Format("2006-01-02 15:04:05 -0700 MST")
////		}
//		fmt.Println(" Nouvelle lune    :\t",nl)
//		fmt.Println(" Premier quartier :\t",pq)
//		fmt.Println(" Pleine lune      :\t",pl)
//		fmt.Println(" Dernier quartier :\t",dq)
//	}
//}

//func jules() {
////	fmt.Println("jules")
////	fmt.Printf("%v %v %v %v %v %v %v\n",t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
////	fmt.Printf("%v %.20f %v\n", t, TimeToJulian(t), JulianToTime(TimeToJulian(t)))
////	y, m, d := Datetime(t)
////	fmt.Printf("%d %d %.20f\n",y, m, d)
////	jd:=julian.CalendarGregorianToJD(y,m,d)
////	fmt.Printf("%.20f\n",jd)
////	y, m, d = julian.JDToCalendar(jd)
////	fmt.Printf("%d %d %.20f\n",y, m, d)
//
//	x:=new(julian.Datetime)
//	fmt.Println(x)
//
////	x=julian.NewDatetimeYMD(1987,1,27.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1987,6,19.5)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1988,1,27.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1988,6,19.5)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1900,1,1.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1600,1,1.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1600,12,31.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(837,4,10.3)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(-1000,7,12.5)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(-1000,2,29.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(-1001,8,17.9)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(-4712,1,1.5)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////
////	x=julian.NewDatetimeYMD(1601,1,0.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1601,13,0.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1601,1,-1.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////	x=julian.NewDatetimeYMD(1601,-1,-1.0)
////	fmt.Println(x,"->",julian.NewDatetimeJD(x.JD()))
////
//	var z *julian.Datetime
////
//	z=julian.NewDatetimeJD(1721119.5)
//	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1954,6,30)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1978,11,14)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1988,4,22)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1583,10,15)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1582,12,31)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
//	z=julian.NewDatetimeYMD(1582,10,15.0)
//	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1582,10,4)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1582,10,3)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
//	z=julian.NewDatetimeYMD(1999,1,1.0)
//	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
//	z=julian.NewDatetimeYMD(2016,1,1.0)
//	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(2000,1,1)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMD(1991,3,31)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1991)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1992)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1993)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1954)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(2000)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1818)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(179)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(711)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeEaster(1243)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMDhmsn(1992,10,13,0,0,0,0)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMDhmsn(2000,1,1,12,0,0,0)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeYMDhmsn(333,2,6,6,0,0,0)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	fmt.Println("year",(z.YearWithFraction()-2000)/100)
////
////	z=julian.NewDatetimeEaster(1991)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeDimancheDePaques(1991)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeLundiDePaques(2016)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeJeudiDelAscension(2016)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
////	z=julian.NewDatetimeLundiDePentecote(2016)
////	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
//
//}
//
//func meeus(a,m,j int) {
//	fmt.Println()
//	z:=julian.NewDatetimeYMD(a,m,float64(j))
//	fmt.Println("z =",z, z.DayOfWeek(), z.DayOfYear(), z.YearWithFraction(), julian.NewDatetimeDoY(z.Year(), z.DayOfYear()))
//}
//
//func paques(a int) {
//	_, m, j := calendrier.Paques(a)
//	fmt.Println("Paques",a,m,j)
//}
//
//func iso(annee, n int) {
//	for y := annee; y < annee+n; y++ {
//		pj := calendrier.PremierJeudiJD(y)
//		a,m,j := calendrier.JD2amj(pj-3)
//		n,s,d := calendrier.AMJ2iso(y,1,1)
//		u,v,w := calendrier.ISO2amj(n,s,d)
//		fmt.Printf("lundi de la premiere semaine de %4d : (%4d-%02d-%02d)\tpremier janvier %4d : iso(%4d-%d-%d) amj(%4d-%02d-%02d))\n",y,a,m,j,y,n,s,d,u,v,w)
//	}
//}

func moy(t []int) float64 {
	var m float64 = 0
	for n := 0; n < len(t); n++ {
		m += (float64(t[n]) - m) / float64(n+1)  
	}
	return m
}

func check(e error) {
    if e != nil {
    	log.Fatal(e)
        panic(e)
    }
}

//func tu() {
////	var x = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,}
////	var y = []int{8,19,0,11,22,3,14,25,6,17,28,9,20,1,12,23,4,15,6,}
////	var xx, xy = make([]int,len(x)), make([]int,len(x))
////	var z []float64 = make([]float64,len(x))
////	for i,v := range x {
////		xx[i] = v*v
////		xy[i] = v*y[i]
////	}
////
////	mx := moy(x)
////	fmt.Println(x," ",mx)
////	my := moy(y)
////	fmt.Println(y," ",my)
////	mxx := moy(xx)
////	fmt.Println(xx," ",mxx)
////	mxy := moy(xy)
////	fmt.Println(xy," ",mxy)
////	a := (mxy - mx*my) / (mxx - mx*mx)
////	b := my - a*mx
////	for i,v := range x {
////		z[i] = a*float64(v) + b
////	}
////
////	fmt.Println("a",a,"b",b)
////	fmt.Println("z",z)
//
//	q := julian.NewDatetimeYMD(2000,1,1.0)
//	fmt.Println(q, q.MJD())
//	q = julian.NewDatetimeYMDhmsn(2016,2,18,19,57,25,0)
//	fmt.Println(q, q.MJD())
//	q = julian.NewDatetimeYMDhmsn(2014,4,10,19,21,0,0)
//	hr, mn, sd, na := q.GMST()
//	fmt.Println(q, q.MJD(), "GMST", hr,"h", mn,"m", sd,"s", na,"n")
//	q = julian.NewDatetimeYMD(2014,4,10.0)
//	fmt.Println(q, q.MJD())
//	annee := 2013
//	easter := julian.NewDatetimeEaster(annee)
//	fmt.Println("annee",annee,"Easter",easter)
//	nowTU := time.Now()
//	nowTD := julian.NewDatetimeTime(nowTU)
//	fmt.Println("nowTU",nowTU,"nowTD",nowTD,nowTD.UTC())
//	q = julian.NewDatetimeYMDhmsn(2013,5,10,0,29,31,0)
//	fmt.Println(q, q.UTC())
//	dt := julian.NewDatetimeYMDhmsn(2013,5,10,0,29,31,0)
//	fmt.Printf("TD=%v TU=%v\n",dt,dt.UTC())
//}
//
//func test_vsop87_ref() {
//	var jds = []float64{2451545.0, 2415020.0, 2378495.0, 2341970.0, 2305445.0, 2268920.0, 2232395.0, 2195870.0, 2159345.0, 2122820.0}
//	w, err := os.Create("output/vsop87_chk.txt")
//	check(err)
//	defer w.Close()
//
//	for _, ivers := range vsop87.VERS {
//		for _, ibody := range vsop87.BODIES {
//			for ijd := 0; ijd < len(jds); ijd++ {
//				jd := jds[ijd]
//				dt := julian.NewDatetimeJD(jd)
//				r, err := vsop87.Vsop87_ref(jd, ivers, ibody)
//				if nil != err {
//					fmt.Println(err)
//				} else {
//					fmt.Fprintf(w," %-8s %-7s     JD%8.1f  %02d/%02d/%4d %2dh TDB\n",ivers,ibody,jd,dt.Day(),dt.Month(),dt.Year(),dt.Hour())
//					fmt.Fprintf(w," %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",r.Name(0),r.Value(0),r.Unit(0) ,r.Name(1),r.Value(1),r.Unit(1) ,r.Name(2),r.Value(2),r.Unit(2))
//					fmt.Fprintf(w," %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",r.Name(3),r.Value(3),r.Unit(3) ,r.Name(4),r.Value(4),r.Unit(4) ,r.Name(5),r.Value(5),r.Unit(5))
//					fmt.Fprintln(w)
//				}
//			}
//			fmt.Fprintln(w)
//		}
//		fmt.Fprintln(w)
//	}
//
//}
//
//func test_vsop87() {
//	var jds = []float64{2451545.0, 2415020.0, 2378495.0, 2341970.0, 2305445.0, 2268920.0, 2232395.0, 2195870.0, 2159345.0, 2122820.0}
//	w, err := os.Create("output/vsop87_chk_ref.txt")
//	check(err)
//	defer w.Close()
//
//	for _, ivers := range vsop87.VERS {
//		for _, ibody := range vsop87.BODIES {
//			for ijd := 0; ijd < len(jds); ijd++ {
//				jd := jds[ijd]
//				dt := julian.NewDatetimeJD(jd)
//				r, err := vsop87.Vsop87(jd, ivers, ibody)
//				if nil != err {
//					fmt.Println(err)
//				} else {
//					fmt.Fprintf(w," %-8s %-7s     JD%8.1f  %02d/%02d/%4d %2dh TDB\n",ivers,ibody,jd,dt.Day(),dt.Month(),dt.Year(),dt.Hour())
//					fmt.Fprintf(w," %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",ivers.CName(0),r[0],ivers.CUnit(0),ivers.CName(1),r[1],ivers.CUnit(1),ivers.CName(2),r[2],ivers.CUnit(2))
//					fmt.Fprintf(w," %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",ivers.CName(3),r[3],ivers.CUnit(3),ivers.CName(4),r[4],ivers.CUnit(4),ivers.CName(5),r[5],ivers.CUnit(5))
//					fmt.Fprintln(w)
//				}
//			}
//			fmt.Fprintln(w)
//		}
//		fmt.Fprintln(w)
//	}
//
//}
//
//func calc_vsop87_ref(jd float64, ivers vsop87.Version, ibody vsop87.Body) {
//	dt := julian.NewDatetimeJD(jd)
//	r, err := vsop87.Vsop87_ref(jd , ivers, ibody)
//	if nil != err {
//		fmt.Println(err)
//	} else {
//		if nil != r {
//			xd := vsop87.Radian2Degre(r.Value(0))
//			d, m, s := vsop87.Radian2DegreMS(r.Value(0))
//			xh := vsop87.Radian2Hour(r.Value(0))
//			hr, mn, sd := vsop87.Radian2HourMS(r.Value(0))
//			fmt.Println(ivers.Name())
//			fmt.Printf(" %-8s %-7s     JD%8.1f  %02d/%02d/%4d %2dh TDB\n",ivers,ibody,jd,dt.Day(),dt.Month(),dt.Year(),dt.Hour())
//			fmt.Printf(" %s    %s    %s\n",r.String(0),r.String(1),r.String(2))
//			fmt.Printf(" %s    %s    %s\n",r.String(3),r.String(4),r.String(5))
//			fmt.Println(xd,d,m,s,xh,hr,mn,sd)
//			fmt.Println(vsop87.Degre2Radian(xd),vsop87.DegreMS2Radian(d, m, s),vsop87.Hour2Radian(xh),vsop87.HourMS2Radian(hr, mn, sd))
//		}
//	}
//}
//
//func calc_vsop87(jd float64, ivers vsop87.Version, ibody vsop87.Body) *[6]float64 {
//	dt := julian.NewDatetimeJD(jd)
//	r, err := vsop87.Vsop87(jd , ivers, ibody)
//	if nil != err {
//		fmt.Println(err)
//		return nil
//	} else {
//		if nil != r {
//			fmt.Println(ivers.Name())
//			fmt.Printf(" %-8s %-7s     JD%8.1f  %02d/%02d/%4d %2dh TDB\n",ivers,ibody,jd,dt.Day(),dt.Month(),dt.Year(),dt.Hour())
//			fmt.Printf(" %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",ivers.CName(0),r[0],ivers.CUnit(0),ivers.CName(1),r[1],ivers.CUnit(1),ivers.CName(2),r[2],ivers.CUnit(2))
//			fmt.Printf(" %-2s %13.10f %-5s    %-2s %13.10f %-5s    %-2s %13.10f %-5s\n",ivers.CName(3),r[3],ivers.CUnit(3),ivers.CName(4),r[4],ivers.CUnit(4),ivers.CName(5),r[5],ivers.CUnit(5))
//		}
//		return r
//	}
//}

func main() {

//	meeus(1582,10,15)
//	meeus(1999,1,1)
//	meeus(2016,2,6)
//	
//	paques(2016)
//	paques(2000)
//	paques(1993)
//	paques(179)
//	paques(711)
//	paques(1243)

//	iso(2010,20)
//	tu()
	
//	test_vsop87()
//	test_vsop87_ref()
	
//	vsop_87()
	
//	dt1 := julian.NewDatetimeYMD(2015,2,15)
//	moon1 := dt1.NextMoon()
//	fmt.Printf("moon1%v %v\n",moon1,moon1.Dt.UTC())
//
//	dt2 := julian.NewDatetimeYMD(2044,1,15)
//	moon2 := dt2.NextMoon2()
//	fmt.Printf("moon2 %v %v\n",moon2,moon2.Dt.UTC())
//
//	dt3 := julian.NewDatetimeYMD(1977,2,15)
//	moon3 := dt3.NextMoon()
//	fmt.Printf("moon3 %v %v\n",moon3,moon3.Dt.UTC())
//	moon32 := dt3.NextMoon2()
//	fmt.Printf("moon3(2)%v %v\n",moon32,moon32.Dt.UTC())
//
//	dt4 := julian.NewDatetimeYMD(2016,1,1)
//	moon4 := dt4.NextMoon()
//	fmt.Printf("moon4 %v %v\n",moon4,moon4.Dt.UTC())

//	annee := 2016
//	for mois := 1; mois <= 4; mois++ {
//		fmt.Printf("\n%d %02d\n",annee,mois)
//		moons := julian.Moons(annee,mois)
//		for _, m := range moons {
//			fmt.Printf(" %v %v\n",m,m.Dt.Local())
//		}
//	}
//	for an := 1972; an <= 2016 ; an++ {
//		for mo := 1; mo < 8; mo += 6 {
//			dt:=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//		}
//	}
//	an := 2020
//	mo := 1
//	dt:=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//
//	an = 1961
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1961
//	mo = 8
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1962
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1963
//	mo = 11
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1964
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1964
//	mo = 4
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1964
//	mo = 9
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1965
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1965
//	mo = 3
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1965
//	mo = 7
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1965
//	mo = 9
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1966
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1968
//	mo = 2
//	dt=julian.NewDatetimeYMD(an,mo,1)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
////fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1972
//	mo = 1
//	dt=julian.NewDatetimeYMD(an,mo,0)
//fmt.Printf("%v %f %.3f %.3f\n",dt,dt.MJD(),julian.DeltaT(an),dt.DeltaTT2UTC())
//fmt.Printf("	var mjd_%04d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())

//	an := 948
//	mo := 1
//	dt := julian.NewDatetimeYMD(an,mo,1)
////fmt.Printf("%v %f %.3f\n",dt,dt.JD(),dt.MJD())
//fmt.Printf("	const mjd_%d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	an = 1620
//	mo = 1
//	dt = julian.NewDatetimeYMD(an,mo,1)
////fmt.Printf("%v %f %.3f\n",dt,dt.JD(),dt.MJD())
//fmt.Printf("	const mjd_%d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())

//	an := 2016
//	mo := 7
//	dt := julian.NewDatetimeYMD(an,mo,1)
////fmt.Printf("%v %f %.3f\n",dt,dt.JD(),dt.MJD())
//fmt.Printf("	const mjd_%d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//
//	an = 2000
//	mo = 1
//	dt = julian.NewDatetimeYMD(an,mo,1)
////fmt.Printf("%v %f %.3f\n",dt,dt.JD(),dt.MJD())
//fmt.Printf("	const mjd_%d_%02d_%02d float64 = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//
//	table := map[int]float64 {
//			1620:121,
//			1630:82,
//			1640:60,
//			1650:46,
//			1660:35,
//			1670:24,
//			1680:14,
//			1690:8,
//			1700:7,
//			1710:9,
//			1720:10,
//			1730:10,
//			1740:11,
//			1750:12,
//			1760:14,
//			1770:15,
//			1780:16,
//			1790:16,
//			1800:13.1,
//			1810:12.0,
//			1820:11.6,
//			1830:7.1,
//			1840:5.4,
//			1850:6.8,
//			1860:7.7,
//			1870:1.4,
//			1880:-5.5,
//			1890:-6.0,
//			1900:-2.8,
//			1910:10.4,
//			1920:21.1,
//			1930:24.0,
//			1940:24.3,
//			1944:26.2,
//			1948:28.2,
//			1952:30.0,
//			1956:31.4,
//			1960:33.1, 
////			1964:35.0, 
////			1968:38.3, 
////			1972:42.2, 
////			1976:46.5, 
////			1980:50.5, 
////			1984:53.8, 
////			1988:55.8, 
////			1992:58.3, 
////			1996:61.6, 
////			2000:63.8, 
////			2004:64.6, 
////			2006:64.9, 
////			2008:65.5, 
////			2010:66.1, 
////			2012:66.6, 
////			2014:67.3, 
//	}
//	fmt.Println("***")
//	for k, _ := range table {
//		dt := julian.NewDatetimeYMD(k,7,1)
//		fmt.Printf("dt %s %v\n",dt.Format(),dt.UTC())
//	}
//	var keys []int
//	keys = make([]int,len(table))
//	i := 0
//	for k, _ := range table {
//		keys[i] = k
//		i++		
//	}
//	sort.Ints(keys)
//	for j := 0; j < len(keys); j++ {
//		k := keys[j]
//		dt := julian.NewDatetimeYMD(k,1,1)
//fmt.Printf("	const mjd_%d_%02d_%02d int = %5.0f //	const jd_%04[1]d_%02[2]d_%02[3]d float64 = %.1[5]f\n",dt.Year(),dt.Month(),dt.Day(),dt.MJD(),dt.JD())
//	}
//	fmt.Println("	table := map[int]float64 {")
//	for j := 0; j < len(keys); j++ {
//		k := keys[j]
//		dt := julian.NewDatetimeYMD(k,1,1)
//fmt.Printf("		mjd_%d_%02d_%02d:%+6.1f,\n",dt.Year(),dt.Month(),dt.Day(),table[k])
////fmt.Printf("	%+6.0f:%+6.1f\n",dt.MJD(), table[k])
//	}
//	fmt.Println("	}")
//	fmt.Println("***")
//fmt.Println("\n******\n")
//	annee := 2016
//fmt.Println("annee", annee)
//	for mois := 1; mois <= 12; mois++ {
//fmt.Println("mois",mois)
//	for _, moon := range julian.Moons(annee,mois) {
//fmt.Printf("moon %v %v\n",moon,moon.Dt.UTC())
//		jdmoon := moon.Dt.JD()
//		dt := correct(jdmoon)
//fmt.Printf("***dt=%f %fs\n",dt,dt*86400)
//		jdmoon += dt
//		emb, _ := vsop87.Vsop87(jdmoon, vsop87.HELIOCENTRIC_RECTANGULAR_COORDINATES_J2000, vsop87.EMB)
//		x0, y0, z0, xv0, yv0, zv0 := emb[0], emb[1], emb[2], emb[3], emb[4], emb[5]
//		l0, b0, r0, lv0, bv0, rv0 := spheric(x0, y0, z0, xv0, yv0, zv0)
//		ear, _ := vsop87.Vsop87(jdmoon, vsop87.HELIOCENTRIC_RECTANGULAR_COORDINATES_J2000, vsop87.EARTH)
//		x1, y1, z1, xv1, yv1, zv1 := ear[0], ear[1], ear[2], ear[3], ear[4], ear[5]
//		l1, b1, r1, lv1, bv1, rv1 := spheric(x1, y1, z1, xv1, yv1, zv1)
//		dx1, dy1, dz1, dxv1, dyv1, dzv1 := ear[0]-emb[0], ear[1]-emb[1], ear[2]-emb[2], ear[3]-emb[3], ear[4]-emb[4], ear[5]-emb[5]
//		// rapport masse lune sur masse terre 
//		const μ = 0.012300034 // IERS 1992
//		dx2, dy2, dz2, dxv2, dyv2, dzv2 := -dx1/μ, -dy1/μ, -dz1/μ, -dxv1/μ, -dyv1/μ, -dzv1/μ
//		x2, y2, z2, xv2, yv2, zv2 := x0+dx2, y0+dy2, z0+dz2, xv0+dxv2, yv0+dyv2, zv0+dzv2
//		l2, b2, r2, lv2, bv2, rv2 := spheric(x2, y2, z2, xv2, yv2, zv2)
//		dx3, dy3, dz3, du3, dv3, dw3 := x2-x1, y2-y1, z2-z1, xv2-xv1, yv2-yv1, zv2-zv1
//
//fmt.Printf("  emb x0=%f, y0=%f, z0=%f, xv0=%f, yv0=%f, zv0=%f\n",x0, y0, z0, xv0, yv0, zv0)
//fmt.Printf("      l0=%f, b0=%f, r0=%f lv0=%f bv0=%f rv0:%f\n",l0*180/math.Pi, b0*180/math.Pi, r0, lv0, bv0, rv0)
//fmt.Printf("earth x1=%f, y1=%f, z1=%f, xv1=%f, yv1=%f, zv1=%f\n",x1, y1, z1, xv1, yv1, zv1)
//fmt.Printf("      l1=%f, b1=%f, r1=%f lv1=%f bv1=%f rv1:%f\n",l1*180/math.Pi, b1*180/math.Pi, r1, lv1, bv1, rv1)
//fmt.Printf(" moon x2=%f, y2=%f, z2=%f, xv2=%f, yv2=%f, zv2=%f\n",x2, y2, z2, xv2, yv2, zv2)
//fmt.Printf("      l2=%f, b2=%f, r2=%f lv2=%f bv2=%f rv2:%f\n",l2*180/math.Pi, b2*180/math.Pi, r2, lv2, bv2, rv2)
//fmt.Printf("earth-emb   dx1=%+f, dy1=%+f, dz1=%+f, dxv1=%+f, dyv1=%+f, dzv1=%+f\n",dx1, dy1, dz1, dxv1, dyv1, dzv1)
//fmt.Printf("            r1=%+f, dz1=%+f, yv1=%+f, dzv1=%+f\n",math.Sqrt(dx1*dx1+dy1*dy1), dz1, math.Sqrt(dxv1*dxv1+dyv1*dyv1), dzv1)
//fmt.Printf(" moon-emb   dx2=%+f, dy2=%+f, dz2=%+f, dxv2=%+f, dyv2=%+f, dzv2=%+f\n",dx2, dy2, dz2, dxv2, dyv2, dzv2)
//fmt.Printf("            r2=%+f, dz2=%+f, yv2=%+f, dzv2=%+f\n",math.Sqrt(dx2*dx2+dy2*dy2), dz2, math.Sqrt(dxv2*dxv2+dyv2*dyv2), dzv2)
//fmt.Printf(" moon-earth dx3=%+f, dy3=%+f, dz3=%+f, du3=%+f, dv3=%+f, dw3=%+f\n",dx3, dy3, dz3, du3, dv3, dw3)
//fmt.Printf("            r3=%+f, dz3=%+f, v3=%+f, dw3=%+f\n",math.Sqrt(dx3*dx3+dy3*dy3), dz3, math.Sqrt(du3*du3+dv3*dv3), dw3)
//		dt = correct(jdmoon)
//fmt.Printf("===dt=%f %fs\n",dt,dt*86400)
//fmt.Println()
//	}
//	}
//	fmt.Printf("L0 = %.12f + %.12f T + %.12f T2\n",vsop87.DegreMS2Radian(279,41,48.04),vsop87.DegreMS2Radian(0,0,129602768.13),vsop87.DegreMS2Radian(0,0,1.089))
//
//	for b := -90.0; b <= 90.0; b += 30.0 {	
//	fmt.Println()
//	for l := 0.0; l < 360.0; l += 30.0 {
//	cb := math.Cos(b*math.Pi/180)
//	sb := math.Sin(b*math.Pi/180)
//	cl := math.Cos(l*math.Pi/180)
//	sl := math.Sin(l*math.Pi/180)
//	rp, bp, lp := spheric(cb*cl,cb*sl,sb)
//	fmt.Printf("b=%5.1f l=%5.1f => r=%5.1f b=%5.1f l=%5.1f\n",b,l,rp,bp*180/math.Pi,lp*180/math.Pi)
//	}
//	}
//var a,m,j int
//var d float64
//var jd int64
//	fmt.Println("---")
//	a,m,j = calendrier.JJ2amj(0)
//	fmt.Println("jj",calendrier.AMJ2jj(a,m,j),a,m,j)
//	a,m,j = calendrier.JG2amj(0)
//	fmt.Println("jg",calendrier.AMJ2jg(a,m,j),a,m,j)
//	a,m,j = calendrier.JJ2amj(-1)
//	fmt.Println("jj",calendrier.AMJ2jj(a,m,j),a,m,j)
//	a,m,j = calendrier.JG2amj(-1)
//	fmt.Println("jg",calendrier.AMJ2jg(a,m,j),a,m,j)
//	a,m,d = calendrier.JGD2amd(-1)
//	fmt.Println("jg+",calendrier.AMD2jgd(a,m,d),a,m,d)
//	
//	fmt.Println("---")
//	a, m, d = 1957,10,4.81
//	fmt.Println("jd",a,m,d,julian.NewDatetimeYMD(a,m,d))
//	jg := calendrier.AMD2jgd(a,m,d)
//	fmt.Printf("jg* %d %d %f %f\n",a,m,d,jg)
//	fmt.Println("jg",calendrier.AMJ2jg(a,m,int(d)))
//	fmt.Println("jj",calendrier.AMJ2jj(a,m,int(d)))
//	a, m, d = calendrier.JGD2amd(jg)
//	fmt.Printf("jg+ %f %d %d %f\n",jg,a,m,d)
//
//	fmt.Println("---")
//	a, m, d = 2016,7,4.25
//	jdd := calendrier.AMD2jdd(a,m,d)
//	fmt.Printf("jd %d %d %f %v %f\n",a,m,d,julian.NewDatetimeYMD(a,m,d),jdd)
//	jg = calendrier.AMD2jgd(a,m,d)
//	fmt.Printf("jg* %d %d %f %f\n",a,m,d,jg)
//	fmt.Println("jg",calendrier.AMJ2jg(a,m,int(d)))
//	fmt.Println("jj",calendrier.AMJ2jj(a,m,int(d)))
//	a, m, d = calendrier.JGD2amd(jg)
//	fmt.Printf("jg+ %f %d %d %f\n",jg,a,m,d)
//	a, m, d = calendrier.JDD2amd(jdd)
//	fmt.Printf("jdd %f %d %d %f\n",jdd,a,m,d)
//
//	fmt.Println("---")
//	a, m, d = -2016,7,4.25
//	jdd = calendrier.AMD2jdd(a,m,d)
//	fmt.Printf("jd %d %d %f %v %f\n",a,m,d,julian.NewDatetimeYMD(a,m,d),jdd)
//	jg = calendrier.AMD2jgd(a,m,d)
//	fmt.Printf("jg* %d %d %f %f\n",a,m,d,jg)
//	jj := calendrier.AMD2jjd(a,m,d)
//	fmt.Printf("jj* %d %d %f %f\n",a,m,d,jj)
////	fmt.Println("jg",calendrier.AMJ2jg(a,m,int(d)))
////	fmt.Println("jj",calendrier.AMJ2jj(a,m,int(d)))
//	a, m, d = calendrier.JGD2amd(jg)
//	fmt.Printf("jg+ %f %d %d %f\n",jg,a,m,d)
//	a, m, d = calendrier.JJD2amd(jj)
//	fmt.Printf("jj+ %f %d %d %f\n",jj,a,m,d)
//	a, m, d = calendrier.JDD2amd(jdd)
//	fmt.Printf("jdd %f %d %d %f\n",jdd,a,m,d)
//
//	fmt.Println("---")
//	a, m, d = -5011,7,4.25
////	fmt.Println("jd",a,m,d,julian.NewDatetimeYMD(a,m,d))
//	jg = calendrier.AMD2jgd(a,m,d)
//	fmt.Printf("jg* %d %d %f %f\n",a,m,d,jg)
//	jj = calendrier.AMD2jjd(a,m,d)
//	fmt.Printf("jj* %d %d %f %f\n",a,m,d,jj)
////	fmt.Println("jg",calendrier.AMJ2jg(a,m,int(d)))
////	fmt.Println("jj",calendrier.AMJ2jj(a,m,int(d)))
//	a, m, d = calendrier.JGD2amd(jg)
//	fmt.Printf("jg+ %f %d %d %f\n",jg,a,m,d)
//	a, m, d = calendrier.JJD2amd(jj)
//	fmt.Printf("jj+ %f %d %d %f\n",jj,a,m,d)
//
//	fmt.Println("---")
//	a,m,d = 333,1,27.5
//	jj = calendrier.AMD2jjd(a,m,d)
//	jdd = calendrier.AMD2jdd(a,m,d)
//	fmt.Printf("jj* %d %d %f %f %f\n",a,m,d,jj,jdd)
//	fmt.Println("jd",julian.NewDatetimeYMDhmsn(a,m,int(d),12))
//	fmt.Println("jd",julian.NewDatetimeYMD(a,m,d))
//	fmt.Println("jj",calendrier.AMJ2jj(a,m,int(d)))
//	a,m,d = calendrier.JJD2amd(jj)
//	fmt.Printf("jd+ %f %d %d %f\n",jj,a,m,d)
//	a, m, d = calendrier.JDD2amd(jdd)
//	fmt.Printf("jdd %f %d %d %f\n",jdd,a,m,d)
//
////	a,m,j = 201,1,1
////	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
////	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
////	a,m,j = 250,1,1
////	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
////	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
////	a,m,j = 300,1,1
////	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
////	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//
//	fmt.Println("---")
//	a,m,j = -4712,1,1
//	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
//	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//	jd = calendrier.AMJ2jd(a,m,j)
//	fmt.Println(a,m,j,"jd",jd, jd%7,(jd-10)%60)
//	fmt.Println(a,m,j,"jd",julian.NewDatetimeYMDhmsn(a,m,j,12))
//
//	fmt.Println("---")
//	a,m,j = -4713,11,24
//	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
//	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//	jd = calendrier.AMJ2jd(a,m,j)
//	fmt.Println(a,m,j,"jd",jd, jd%7,(jd-10)%60)
////	fmt.Println(a,m,j,"jd",julian.NewDatetimeYMDhmsn(a,m,j,0))
//
//	fmt.Println("---")
//	a,m,j = 1,1,1
//	fmt.Println("***")
//	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
//	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//	fmt.Printf("jgd %d %d %d %f\n",a,m,0,calendrier.AMD2jgd(a,m,0))
//	jd = calendrier.AMJ2jd(a,m,j)
//	fmt.Println(a,m,j,"jd",jd, jd%7,(jd-10)%60)
//	fmt.Println(a,m,j,"jd",julian.NewDatetimeYMDhmsn(a,m,j,0))
//	fmt.Println("***")
//
//	fmt.Println("---")
//	a,m,j = 1582,10,4
//	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
//	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//	jd = calendrier.AMJ2jd(a,m,j)
//	fmt.Println(a,m,j,"jd",jd, jd%7,(jd-10)%60)
//
//	fmt.Println("---")
//	a,m,j = 1582,10,15
//	fmt.Println(a,m,j,"jj",calendrier.AMJ2jj(a,m,j))
//	fmt.Println(a,m,j,"jg",calendrier.AMJ2jg(a,m,j))
//	jd = calendrier.AMJ2jd(a,m,j)
//	fmt.Println(a,m,j,"jd",jd, jd%7,(jd-10)%60)
//	
////	fmt.Println("===")
////	vsop87.Generate_test()
////	fmt.Println("===")
//
////	var err error
//	pluton()
//fmt.Println("==================")
//	series()
//fmt.Println("==================")
//vsop87.Generate_data()
//fmt.Println("==================")
//	fmt.Println("dt",julian.NewDatetimeJD(77432.5))
//	fmt.Println("dt",julian.NewDatetimeJD(625307.5))
//	fmt.Println("dt",julian.NewDatetimeJD(1173182.5))
//	fmt.Println("dt",julian.NewDatetimeJD(1721057.5))
//	fmt.Println("dt",julian.NewDatetimeJD(2268932.5))
//	fmt.Println("dt",julian.NewDatetimeJD(2816818.5))
//	vsop2013.Generate_ref()
//	vsop2013.Compute_ref()
//	vsop2013_test()
//	vsop2013.Generate_db()
//	vsop2013.Generate_sql()

//	json_config()
//	vsop_87()
//	vsop_2013()
//	compare_vsop_87_2013()
//	inpop()
	
	//de432()
}

//func de432() {
//
////	jpl.Jpl_generatedb()
//
//	DENUM, _ := jpl.Jpl_get_const("DENUM")
//	fmt.Println("DENUM",DENUM)
//	LENUM, _ := jpl.Jpl_get_const("LENUM")
//	fmt.Println("LENUM",LENUM)
//	AU, _ := jpl.Jpl_get_const("AU")
//	fmt.Println("AU",AU)
//	JDEPOC, _ := jpl.Jpl_get_const("JDEPOC")
//	depoch := julian.NewDatetimeJD(JDEPOC)
//	fmt.Println("DEPOCH",depoch)
//	CENTER, _ := jpl.Jpl_get_const("CENTER")
//	fmt.Println("CENTER",CENTER)
//	CLIGHT, _ := jpl.Jpl_get_const("CLIGHT")
//	fmt.Println("CLIGHT",CLIGHT)
//	EMRAT, _ := jpl.Jpl_get_const("EMRAT")
//	fmt.Println("EMRAT",EMRAT)
//
//	var by jpl.Jpl_body
//	var jd float64
//
////	jd, by = 2442352.5, jpl.MARS
////	jd, by = 2442413.5, jpl.Jpl_body(7)
////	jd, by = 2434529.5, jpl.Jpl_body(6)
////	jd, by = 2434986.5, jpl.Jpl_body(9)
////	jd, by = 2441408.5, jpl.Jpl_body(1)
////	jd, by = 2448196.5, jpl.Jpl_body(1)	// 2441683.5, jpl.Jpl_body(1)
//	jd, by = 2442413.5, jpl.Jpl_body(7)
////	jd, by = 2437331.5, jpl.Jpl_body(10)
////	jd, by = 2437331.5, jpl.MOON
////	jd, by = 2438364.5, jpl.EARTH
////	jd, by = 2453948.5, jpl.EMB
////	jd, by = 2436112.5, jpl.EARTH
////	jd, by = 2437116.5, jpl.SUN
//	dte := julian.NewDatetimeJD(jd)
//	fmt.Println("jd",dte)
//	pv, err := jpl.Jpl_eph(by,jd)
//	if nil != err {
//		fmt.Println("err",err)
//	} else {
//		fmt.Printf("%v	%.1f	pv	",by,jd)
//		for i := 0; i < 6; i++ {
//			fmt.Printf("%.16e	",pv[i])
//		}
//		fmt.Println()
//		pvc := vsop2013.HEQ2hec(pv)
//		fmt.Printf("pvc %v	%.1f	pv	\n",by,jd)
//		for i := 0; i < 3; i++ {
//			fmt.Printf("%.10f	",pvc[i])
//		}
//		fmt.Println()
//		for i := 3; i < 6; i++ {
//			fmt.Printf("%.10f	",pvc[i])
//		}
//		fmt.Println()
//		pv6 := calc_vsop87(jd, vsop87.BARYCENTRIC_RECTANGULAR_COORDINATES_J2000, vsop87.Body(by))
//		for i := 0; i < 6; i++ {
//			fmt.Printf("%e\n",pv6[i]-pvc[i])
//		}
//	}
//}

func inpop() {
//	wd, err := os.Getwd()
//	fmt.Println(os.Args[0],wd,err)
//	res := inpop13c.ReadHeader(dir+"inpop13c_TCB_m1000_p1000_bigendian.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TCB_m100_p100_littleendian.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TCB_m100_p100_tcg.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TCB_m1000_p1000_littleendian.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TCB_m1000_p1000_tcg.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TDB_m100_p100_littleendian.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TDB_m100_p100_tt.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TDB_m1000_p1000_littleendian.dat")
//	res := inpop13c.ReadHeader(dir+"inpop13c_TDB_m1000_p1000_tt.dat")
//	res := inpop13c.ReadHeaders(configuration.Configuration.InputDir+"/inpop13c/example1.dat")
//	fmt.Printf("res=%v\n",res)
//	for _,c := range []string{"UNITE",
//							"FORMAT",
//							"TIMESC",
////							"Z_Mar ",
////							"YD_Ven",
////							"Z_Ven",
////							"Y_EMB",
////							"Z_LIBM",
////							"RMOON",
////							"CLIGHT",
////							"S33M",
////							"GM_Ven",
//	} {
//		f,_ := res.GetConst(c)
//		fmt.Printf("'%-6s' = %+.16e\n",c, f)
//	}
	
}

//func json_config() {
//	var config = configuration.Config{
//		InputDir:"/Users/philippe/Projects/workspace-java/GoMeeusTest/input",
//		OutputDir:"/Users/philippe/Projects/workspace-java/GoMeeusTest/output",
//		TchebNmax:15,
//		SQL:configuration.SqlConfig{
//			Host:"127.0.0.1",
//			Port:3306,
//			User:"test",
//			Pwd:"test",
//			Db:"test",
//		},
//	}
//fmt.Println()
//	enc := json.NewEncoder(os.Stdout)
//	enc.Encode(config)
//fmt.Println()
//}

//func compare_vsop_87_2013() {
//	for _,jd := range []float64{2451545.0/*,1721057.5,2268932.5,2816818.5*/} {
//		fmt.Printf("jd %f %v\n",jd,julian.NewDatetimeJD(jd))
//		dt := julian.NewDatetimeJD(jd)
//		for _, p := range []int{/*1,2,*/3/*,4,5,6,7,8*/} {
//			fmt.Println()
//			r2013, err := vsop2013.EphsVsop2013_sql[vsop2013.Planet(p)](jd)
//			if nil != err {
//				fmt.Println(err)
//				continue
//			}
//			r13 := r2013.E()
//			fmt.Println("VSOP2013")
//			fmt.Println("Heliocentric rectangular coordinates J2000")
//			fmt.Printf(" %-8s %-9s   JD%8.1f  %02d/%02d/%4d %2dh TDB\n","VSOP2013",vsop2013.Planet(p),jd,dt.Day(),dt.Month(),dt.Year(),dt.Hour())
//			fmt.Printf(" X  %13.10f au       Y  %13.10f au       Z  %13.10f au   \n",r13[0],r13[1],r13[2])
//			fmt.Printf(" X' %13.10f au/d     Y' %13.10f au/d     Z' %13.10f au/d \n",r13[3],r13[4],r13[5])
//
//			fmt.Println("VSOP87")
//			b := vsop87.Body(p)
//			if 3 == p {
//				b = vsop87.EMB
//			}
//			r87 := calc_vsop87(jd, vsop87.HELIOCENTRIC_RECTANGULAR_COORDINATES_J2000, b)
//			d := 0.0
//			for i := 0; i < 6; i++ {
//				e := r13[i] - r87[i]
//				if d < e {
//					d = e
//				} else if d < -e {
//					d = -e
//				}
//			}
//			fmt.Printf(" %.0e\n",d)
//		}
//	}
//}
//
//func vsop_87() {
//	for _,jd := range []float64{2442457.5/*1721057.5,2268932.5,2816818.5*/} {
//		fmt.Printf("jd %f\n",jd)
//		for _, vers := range []vsop87.Version{
////			vsop87.HELIOCENTRIC_SPHERICAL_COORDINATES_J2000,
//			vsop87.HELIOCENTRIC_RECTANGULAR_COORDINATES_J2000,
//			vsop87.BARYCENTRIC_RECTANGULAR_COORDINATES_J2000,
//		} {
//			for _, body := range []vsop87.Body{
//					vsop87.MARS,
////					vsop87.MERCURY,
////					vsop87.VENUS,
//			} {
//				calc_vsop87(jd, vers, body)
////				calc_vsop87_ref(jd, vers, body)
////				r := calc_vsop87(jd, vers, body)
////				d, m, s := vsop87.Radian2DegreMS(r[0])
////				xd := vsop87.Radian2Degre(r[0])
////				xh := vsop87.Radian2Hour(r[0])
////				hr, mn, sd := vsop87.Radian2HourMS(r[0])
////				fmt.Println(xd,d,m,s,xh,hr,mn,sd)
////				fmt.Println(vsop87.Degre2Radian(xd),vsop87.DegreMS2Radian(d, m, s),vsop87.Hour2Radian(xh),vsop87.HourMS2Radian(hr, mn, sd))
////				fmt.Println(vsop87.Radian2Degre(r[0]),vsop87.Radian2Degre(r[1]))
////
////				calc_vsop87_ref(jd, vers, vsop87.EARTH)
////				r = calc_vsop87(jd, vers, vsop87.EARTH)
////				fmt.Println(vsop87.Degre(r[0]),vsop87.Degre(r[1]))
////				calc_vsop87_ref(jd, ivers, vsop87.EARTH)
////
////				err := vsop87.Generate()
////				if nil != err {
////					fmt.Println(err)
////				}
////
////				err := vsop87.Generate_ref()
////				if nil != err {
////					fmt.Println(err)
////				}
//			}
//		}
//	}
//
//}

//func vsop_2013() {
//	for _,jd := range []float64{1721057.5,2268932.5,2816818.5} {
//		for _, p := range []vsop2013.Planet{vsop2013.MERCURY,vsop2013.VENUS,} {
//			if r, err := vsop2013.EphsVsop2013_sql[p](jd); nil != err {
//				fmt.Printf("[%s %v]\n",p,err)
//			} else {
//				fmt.Printf("%v\n",r)
//			}
//			if r, err := vsop2013.EphVsop2013_db(p,jd); nil != err {
//				fmt.Printf("[%s %v]\n",p,err)
//			} else {
//				fmt.Printf("%v\n",r)
//			}
//		}
//	}
//
//}

func print_eph_ref(r map[float64]map[vsop2013.Planet]*[6]float64) {
	fmt.Println("print_eph_ref")

	var jdkeys []float64
	var pkeys []int

//	for jd, u := range r {
//		fmt.Printf("jd: %.1f\n",jd)
//		for p := range u {
//			pkeys = append(pkeys, int(p))
//		}
//		sort.Ints(pkeys)
//		for p, v := range u {
//			fmt.Printf("	%v r=%v\n",p,*v)
//		}
//	}
//	fmt.Println("----")

	for jd := range r {
		jdkeys = append(jdkeys,jd)
	}
	sort.Float64s(jdkeys)
	for i := 0; i < len(jdkeys); i++ {
		jd := jdkeys[i]
		u := r[jd]
		fmt.Printf("jd: %.1f\n",jd)
		pkeys = pkeys[0:0]
		for p := range u {
			pkeys = append(pkeys, int(p))
		}
		sort.Ints(pkeys)
		for k := 0; k < len(pkeys); k++ {
			p := vsop2013.Planet(pkeys[k])
			v := u[p]
			fmt.Printf("	%9v r=%v\n",p,*v)
		}
	}
}

func print_eph(r map[vsop2013.Planet]map[float64]*[6]float64) {
fmt.Println("print_eph")

	var jdkeys []float64
	var pkeys []int

//	for p, u := range r {
//		fmt.Printf("%v\n",p)
//		for jd, v := range u {
//			fmt.Printf("	JD%.1f r=%v\n",jd,*v)
//		}
//	}
//	fmt.Println("--------")

	for p := range r {
		pkeys = append(pkeys, int(p))
	}
	sort.Ints(pkeys)
	for i := 0; i < len(pkeys); i++ {
		p := vsop2013.Planet(pkeys[i])
		u := r[p]
		fmt.Printf("%v\n",p)
		jdkeys = jdkeys[0:0]
		for jd := range u {
			jdkeys = append(jdkeys,jd)
		}
		sort.Float64s(jdkeys)
		for k := 0; k < len(jdkeys); k++ {
			jd := jdkeys[k]
			v := u[jd]
			fmt.Printf("	JD%.1f r=%v\n",jd,*v)
		}
	}

}

func vsop2013_calc() {
	jd := 2451545.0
	var r87 *[6]float64
	for p := 1; p <= 9; p++ {
		//if p < 9 {
		//	r87, _ = vsop87.Vsop87(jd,vsop87.HELIOCENTRIC_RECTANGULAR_COORDINATES_J2000, vsop87.Body(p))
		//} else {
		//	r87, _ = pluto95.Pluto95(jd)
		//}
		//fmt.Println("r87",r87)
		r2013, _ := vsop2013.EphVsop2013b(vsop2013.Planet(p),jd)
		fmt.Println("r2013",r2013)
		var e [6]float64
		d := 0e0
		for i := 0; i < 6; i++ {
			e[i] = r2013[i] - r87[i]
			if e[i] > d {
				d = e[i]
			} else if -e[i] > d {
				d = -e[i]
			}
		}
		fmt.Printf("%v: d %.0e e %v\n",vsop2013.Planet(p),d,e)
		fmt.Println("---")
	}
}

func vsop2013_test() {
	zeros := []float64{
//		500.0,
		77432.5e0,     // -4500/01/01 0h
		625307.5e0,    // -3000/01/01 0h
		1173182.5e0,   // -1500/01/01 0h
		1721057.5e0,   //  0000/01/01 0h
		2268932.5e0,   //  1500/01/01 0h
		2816818.5e0,   //  3000/02/01 0h
//		5000000.0,
	}
	dat0 := 0	
//	planets := []vsop2013.Planet{}
//	planets := []vsop2013.Planet{vsop2013.MERCURY,
//		vsop2013.PLUTO,
//	}
//	ndat := 2
	planets := vsop2013.PLANETS
	ndat := 5
	step := 136798.e0
	for i := 0; i < len(zeros); i++ {
		jds := make([]float64,ndat)
		for n := 0; n < ndat; n++ {
			jds[n] = zeros[i] +float64(dat0+n)*step
		}
//		ref, err_ref := vsop2013.EphVsop2013_ref(jds,planets)
//		if nil != err_ref {
//			fmt.Println("err",err_ref)
//		}
//		fmt.Println("===================")
//		print_eph_ref(ref)
//		fmt.Println("----------------")
		r, err := vsop2013.EphVsop2013(planets,jds)
		if nil != err {
			fmt.Println("err",err)
		}
		print_eph(r)
//		fmt.Printf("planets:[]vsop2013.Planet{")
//		for _, p := range planets {
//			fmt.Printf("%s,",p)
//		}
//		fmt.Println("}")
//		fmt.Printf("planets:%#v\n",planets)
//		fmt.Printf("jds:[]float64{")
//		for _, jd := range jds {
//			fmt.Printf("%.1f,",jd)
//		}
//		fmt.Println("}")
//		fmt.Printf("jds:%#v\n",jds)
//		fmt.Printf("eph:%#v\n",r)
		fmt.Println("===================")
	}
}

func vsop2013_binfile() {
	vsop2013.Vsop2013_binfile()
	fmt.Println("==================")
}

//func moon() {
//	var a, m, j int
//	var xyz *elp82b.ELP82B_t
//	var jdf *julian.Datetime
//
//	fmt.Println("---")
//	a,m,j = 2047,10,17
//	jdf = julian.NewDatetimeYMDhmsn(a,m,j,0)
//	fmt.Println(a,m,j,"jdf",jdf)
////	xyz, err = elp82b.Elp82b_ref(jdf.JD())
////	if nil != err {
////		fmt.Println("err",err)
////	} else {
////		fmt.Printf("xyz %v\n",xyz)
////	}
//	xyz = elp82b.Elp82b(jdf.JD())
//	fmt.Printf("xyz %v\n",xyz)
//
//	fmt.Println("---")
//	a,m,j = 1993,1,13
//	jdf = julian.NewDatetimeYMDhmsn(a,m,j,0)
//	fmt.Println(a,m,j,"jdf",jdf)
////	xyz, err = elp82b.Elp82b_ref(jdf.JD())
////	if nil != err {
////		fmt.Println("err",err)
////	} else {
////		fmt.Printf("xyz %v\n",xyz)
////	}
//	xyz = elp82b.Elp82b(jdf.JD())
//	fmt.Printf("xyz %v\n",xyz)
//
//	fmt.Println("---")
//	a,m,j = 1938,4,12
//	jdf = julian.NewDatetimeYMDhmsn(a,m,j,0)
//	fmt.Println(a,m,j,"jdf",jdf)
////	xyz, err = elp82b.Elp82b_ref(jdf.JD())
////	if nil != err {
////		fmt.Println("err",err)
////	} else {
////		fmt.Printf("xyz %v\n",xyz)
////	}
//	xyz = elp82b.Elp82b(jdf.JD())
//	fmt.Printf("xyz %v\n",xyz)
//
//	fmt.Println("---")
//	a,m,j = 1883,7,9
//	jdf = julian.NewDatetimeYMDhmsn(a,m,j,0)
//	fmt.Println(a,m,j,"jdf",jdf)
////	xyz, err = elp82b.Elp82b_ref(jdf.JD())
////	if nil != err {
////		fmt.Println("err",err)
////	} else {
////		fmt.Printf("xyz %v\n",xyz)
////	}
//	xyz = elp82b.Elp82b(jdf.JD())
//	fmt.Printf("xyz %v\n",xyz)
//
//	fmt.Println("---")
//	a,m,j = 1828,10,5
//	jdf = julian.NewDatetimeYMDhmsn(a,m,j,0)
//	fmt.Println(a,m,j,"jdf",jdf)
////	xyz, err = elp82b.Elp82b_ref(jdf.JD())
////	if nil != err {
////		fmt.Println("err",err)
////	} else {
////		fmt.Printf("xyz %v\n",xyz)
////	}
//	xyz = elp82b.Elp82b(jdf.JD())
//	fmt.Printf("xyz %v\n",xyz)
//
////	elp82b.Generate()
//
//	fmt.Println("==================")
//}

//func pluton() {
//	var tjds []float64 = []float64{
//		2451545.0,
//		2451515.0,
//		2451485.0,
//		2451455.0,
//		2451425.0,
//		2451395.0,
//		2451365.0,
//		2451335.0,
//		2451305.0,
//		2451275.0,
//		2451245.0,
//		2451215.0,
//		2451185.0,
//		2451155.0,
//		2451125.0,
//		2451095.0,
//		2451065.0,
//		2451035.0,
//		2451005.0,
//		2450975.0,
//		2450945.0,
//		2450915.0,
//		2450885.0,
//		2450855.0,
//		2450825.0,
//		2450795.0,
//		2450765.0,
//		2450735.0,
//		2450705.0,
//		2450675.0,
//		2450645.0,
//		2450615.0,
//		2450585.0,
//		2450555.0,
//		2450525.0,
//		2450495.0,
//	}
//
//	for _, tjd := range tjds {
//	r, _ := pluto95.Pluto95(tjd)
///*
//    Rectangular Coordinates of Pluto at Date: JD2451545.00000000
//  X= -9.87600250590294  Y=-27.97918968704569  Z= -5.75368219264631 au
//  X.=  .00302869050594  Y.= -.00112766304604  Z.= -.00126513753728 au/day
//*/
//	fmt.Println()
//	fmt.Printf("    Rectangular Coordinates of Pluto at Date: JD%.1f\n",tjd)
//	fmt.Printf("  X=%18.14f  Y=%18.14f  Z=%18.14f au\n",r[0],r[1],r[2])
//	fmt.Printf("  X.=%17.14f  Y.=%17.14f  Z.=%17.14f au/d\n",r[3],r[4],r[5])
//
//	}
//}

//func series() {
////	series96.Generate_ref()
////	var tjds []float64 = []float64{
////		2415020.5,
////		2429520.75,
////		2444021.0,
////		2458521.25,
////		2473021.5,
////	}
////
////	for _, planet := range series96.PLANETS {
//////	for ipla := 1; ipla <= 1; ipla ++ {
//////		planet := series96.Planet(ipla)
////		for _, tjd := range tjds {
////			dt := julian.NewDatetimeJD(tjd)
////			r, err := series96.Series_ref(tjd, planet)
////			if nil != err {
////				fmt.Printf("err=%v\n",err)
////			} else {
////				fmt.Println()
////				fmt.Printf("%s\n",planet)
////				fmt.Printf("JD: %15.5f  Date:%s\n",tjd,dt)
////				fmt.Printf("x : %15.12f  y : %15.12f  z : %15.12f au\n",r[0],r[1],r[2])
////				fmt.Printf("x': %15.12f  y': %15.12f  z': %15.12f au/d\n",r[3],r[4],r[5])
////			}
////			break
////		}
////		break
////	}
//
////1001  format (////
////     .        2x,'--------------------------------------------'/
////     .        2x,'PLANEPH : ASTROMETRIC COORDINATES OF PLANETS'/
////     .        2x,'--------------------------------------------'/)
//
//
////1002  format (2x,a8,2x,'JD',f13.5,
////     .        3x,'alpha: ',f12.9,' rad',
////     .        3x,'delta: ',f12.9,' rad')
//
////1003  format (//2x,'Continue : Press Enter'//)
//
////1004  format (//2x,'Program Ended : Press Enter'//)
//
////*     ------------
////*     Declarations
////*     ------------
//
////	dimension v1(6),v2(6),w(3)
//	const iv10, iv20, iw0 = 1, 1, 1
//	var v1, v2 *[6]float64
//	var w [3]float64
//	var err error
//
////      data tlight/0.577551831d-2/
////      data dpi/6.283185307179586d0/
// 	const (
// 		tlight = 0.577551831e-2
// 		dpi = 6.283185307179586e0
// 	)
//
////      data pla/'Mercury','Venus','*','Mars','Jupiter','Saturn','Uranus','Neptune','Pluto'/
//	const ipla0 = 1
//	var pla = [...]string{"Mercury","Venus","*","Mars","Jupiter","Saturn","Uranus","Neptune","Pluto"}
//
////*     ------------
////*     Dates limits
////*     ------------
////*
////      data t1/2415020.5d0/
////      data t2/2487980.5d0/
//	const (
//		t1 = 2415020.5e0
//		t2 = 2487980.5e0
//	)
//
////*     -------------------
////*     Initial julian date
////*     -------------------
////*
////*     From : JD2415020.5d0 (1 Jan 1900 0h)
////*     To   : JD2487980.5d0 (4 Oct 2099 0h)
////*
////      t0=2415020.5d0
//	t0 := 2415020.5e0
//
////*     ---------------
////*     Number of dates
////*     ---------------
////
////      ndat=5
//	ndat := 1
//
////*     ------------------
////*     Tab interval (day)
////*     ------------------
////*
////      step=14500.25d0
//
//	step := 14500.25e0
//
////*     ----------
////*     Test dates
////*     ----------
////*
////      tlim=t0+(ndat-1)*step
////      if (t0.lt.t1.or.t0.gt.t2) stop '  Error : Date'
//	tlim := t0+float64(ndat-1)*step
//	if t0 < t1 || t0 > t2 {
//		panic("  Error : Date")
//	}
//	if tlim < t1 || tlim > t2 {
//		panic("  Error : Date")
//	}
//
////*     ------------
////*     Planet index
////*     ------------
////*
////*     ipla=1   Mercury
////*     ipla=2   Venus
////*     ipla=4   Mars
////*     ipla=5   Jupiter
////*     ipla=6   Saturn
////*     ipla=7   Uranus
////*     ipla=8   Neptune
////*     ipla=9   Pluto
////*
////      ipla=1
//	const ipla_first = 1
//	const ipla_last = 9
//	for ipla := ipla_first; ipla <= ipla_last; ipla++ {
//
//		if 3 == ipla {
//			continue
//		}
//
//		planet := series96.Planet(ipla)
//		fmt.Println("planet",planet)
//
////*     -----------------
////*     Test planet index
////*     -----------------
////*
////      if (ipla.lt.1.or.ipla.eq.3.or.ipla.gt.9) stop '  Error : Planet'
//		if ipla < 1 || ipla == 3 || ipla > 9 {
//			panic("  Error : Planet")
//		}
//
////*     ----------
////*     Dates loop
////*     ----------
////*
////      do n=1,ndat
////*
////         if (mod(n,10).eq.1) write (*,1001)
//		for n := 1; n <= ndat; n++ {
//			if n%10 == 1 {
//				fmt.Println("--------------------------------------------\nPLANEPH : ASTROMETRIC COORDINATES OF PLANETS\n--------------------------------------------\n")
//			}
////         tjd=t0+(n-1)*step
//			tjd := t0+float64(n-1)*step
//
////*        -------------------------------
////*        Compute astrometric coordinates
////*        -------------------------------
////*
////         call SERIES (tjd,3,luemb,v1,ierr)
////         if (ierr.ne.0) stop '  Error : EMB series'
////         call EARTH (tjd,v2)
////         do i=1,3
////            v1(i)=v1(i)-v2(i)
////         enddo
//			v1, err = series96.Series(tjd, series96.EMB)
//			if nil != err {
//				panic("  Error : EMB series")
//			}
////fmt.Printf("v1=%.16f %.16f %.16f %.16f %.16f %.16f\n",v1[0],v1[1],v1[2],v1[3],v1[4],v1[5])
//
//			w = series96.Earth(tjd)
////fmt.Printf("earth: w=%.16f %.16f %.16f\n",w[0],w[1],w[2])
//			for i := 0; i < 3; i++ {
//				v1[i]-=w[i]
//			}
//
////         call SERIES (tjd,ipla,lupla,v2,ierr)
////         if (ierr.ne.0) stop '  Error : planetary series'
//			v2, err = series96.Series(tjd, planet)
//			if nil != err {
//				panic("  Error : planetary series")
//			}
//			fmt.Printf("%s\n",planet)
//			fmt.Printf("JD: %15.5f\n",tjd)
//			fmt.Printf("x : %15.12f  y : %15.12f  z : %15.12f au\n",v2[0],v2[1],v2[2])
//			fmt.Printf("x': %15.12f  y': %15.12f  z': %15.12f au/d\n",v2[3],v2[4],v2[5])
//
////         do k=1,2
////            do i=1,3
////               w(i)=v2(i)-v1(i)
////            enddo
////            dist=dsqrt(w(1)*w(1)+w(2)*w(2)+w(3)*w(3))
////            t=tjd-dist*tlight
////            call SERIES (t,ipla,lupla,v2,ierr)
////            if (ierr.ne.0) stop '  Error : planetary series'
////         enddo
//
////	Light correction
//			for k := 1; k <= 2; k++ {
//				for i := 0; i < 3; i++ {
//					w[i]=v2[i]-v1[i]
//				}
//				dist := dsqrt(w[0]*w[0]+w[1]*w[1]+w[2]*w[2])
//				t := tjd-dist*tlight
//				v2, err = series96.Series(t,series96.Planet(ipla))
//				if nil != err {
//					panic("  Error : planetary series")
//				}
//	     	}
////fmt.Printf("v1=%.16f %.16f %.16f %.16f %.16f %.16f\n",v1[0],v1[1],v1[2],v1[3],v1[4],v1[5])
////fmt.Printf("v2=%.16f %.16f %.16f %.16f %.16f %.16f\n",v2[0],v2[1],v2[2],v2[3],v2[4],v2[5])
//
////         do i=1,3
////            w(i)=v2(i)-v1(i)
////         enddo
//			for i := 0; i < 3; i++ {
//				w[i] = v2[i] - v1[i]
//			}
//
////*        ---------------
////*        Display results
////*        ---------------
////*
////         alpha=datan2(w(2),w(1))
////         if (alpha.lt.0.0) alpha=alpha+dpi
////         p=dsqrt(w(1)*w(1)+w(2)*w(2))
////         delta=datan(w(3)/p)
////*
////         write (*,1002) pla(ipla),tjd,alpha,delta
////*
////         if (mod(ndat,10).eq.1) then
////            write (*,1003)
////            read (*,*)
////         endif
////*
////      enddo
//
////fmt.Printf("w=%.16f %.16f %.16f\n",w[0],w[1],w[2])
//			alpha := datan2(w[1],w[0])
////fmt.Printf("alpha=%v\n",alpha)
//			if (alpha < 0.0) {
//				alpha+=dpi
//			}
//			p := dsqrt(w[0]*w[0]+w[1]*w[1])
//			delta := datan(w[2]/p)
//
//			fmt.Printf("%-8s  JD%13.5f   alpha: %12.9f rad   delta: %12.9f rad\n",pla[ipla-ipla0],tjd,alpha,delta)
//
//			for i := 0; i < 5; i++ {
//				v2[i] -= v1[i]
//			}
//			v3 := series96.XYZ2adr(v2)
//			fmt.Println(v3)
//			fmt.Println()
//		}
//	}
//
//}

func dsqrt(x float64) float64 {
	return math.Sqrt(x)
}

func datan2(y, x float64) float64 {
	return math.Atan2(y, x)
}

func datan(x float64) float64 {
	return math.Atan(x)
}