package main;
import (
	"math"

)


// this helps me with finding data or differnce of two points in km's after i multiply by 1000 its now in meter  
func haversineDistance(lat1,long1,lat2,long2 float64)(float64){
	const R = 6371; 

    dLat := (lat2 - lat1) * math.Pi / 180.0; 
    dLon := (long2 - long1) * math.Pi / 180.0; 

	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c * 1000
}




// this will help me to get if the user is under the variable meter i am passing and use haverstine distance to calculate it
func inRadiuscheck(lat1,long1,lat2,long2 float64, variableMeter float64)(bool){
  result:=haversineDistance(lat1,lat2,long1,long2); 
  return result<=variableMeter
}


