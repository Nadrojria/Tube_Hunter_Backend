package model

type Location struct {
    ID				int    		`json:"id"`
    Country    		string 		`json:"country"`
    City     		string 		`json:"city"`
    Lat    	        int 		`json:"lat"`
    Long    	    int 		`json:"long"`
}