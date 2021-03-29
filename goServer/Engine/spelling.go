package Engine;

import (
	"bytes"
	"sort"
	"strings"
)

func compareWords(input,word string) float64{
	var input_arr []byte = []byte(input)
	var word_arr []byte = []byte(word)

	var require float64 = float64(len(input_arr));

	var found float64;

	for n,i := range input_arr{

		var letter_idx = bytes.IndexAny(word_arr,string(i))
		if letter_idx == n {
			found = found + 1.0
			found = found + found * 0.0001
		} else if letter_idx != -1 {

		 	if letter_idx == n+1 || letter_idx == n-1 {
		 		found = found + 0.8
		 		found = found + found * 0.0001
		 	} else if letter_idx == n+2 || letter_idx == n-2 {
		 		found = found + 0.5
		 		found = found + found * 0.0001
		 	} else {
		 		found = found + 0.2
		 		found = found + found * 0.0001
		 	}
		}
	}
	var per float64 = float64( (found / require) * 100.0)
	
	return per 

}


func sortWords(input string) [] string{
	var first_letter string = strings.Split(input,"")[0]

	var result_map = make(map[float64]string,1)
	for _,i := range wordsHolder{
		if strings.Index(i , first_letter) == 0 {

			per := compareWords(input,i)
			if per == 0.0 || per < 50{continue} 

			result_map[per] = i;
		}
	}
	var keys []float64 = make([]float64,1)
	for k,_ := range result_map{
		keys = append(keys,k)
	}
	sort.Float64s(keys);

	 var result []string;
	for _,i := range keys{
		var word = result_map[i];
		if word == ""{continue}
		result = append(result,word)
	}
	

	for i,j := 0, len(result) - 1 ; i<j ;i,j = i+1 , j-1{
		result[i] , result[j] = result[j] ,result[i]
	}

	return result
}