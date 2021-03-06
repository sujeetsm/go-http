package webservice

import (
	"net/http"
	"log"
	"encoding/json"
	"net/url"
)

/*
 * Process incoming get request
 * POST/PUT request are not allowed
*/

func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	
	log.Println("Incoming Request:", r.Method)
	switch r.Method {
	case http.MethodGet:
		Get(w, r)
		break
	default:
		Error(&w, 405, "Method not allowed", "Method not allowed", nil)
		break
	}
}

/*
 * Func to fetch the names from uinames.com/api
 * Then pass the names struct to fetch the joke
*/

func Get(w http.ResponseWriter, r *http.Request) {
	var names_data UINames

	names_response, err := http.Get("http://uinames.com/api")
	if err != nil {
		log.Println("Error :",err)
		Error(&w, 408, "Unable to fetch Names", "destination unreachable", nil)
	} else {
		err = json.NewDecoder(names_response.Body).Decode( &names_data)
		if err != nil {
			Error(&w, 415, "Unable to Decode Names", "MalFormed response", nil)
		} else {
			names_data.GetJoke(w,r)
			
		}
	}
}

/*
 * Func to fetch the jokes based of UINames reciever.
 * Builds/Encodes the URL for the request to get jokes
 * 
*/

func (uinames *UINames) GetJoke(w http.ResponseWriter, r *http.Request) {
	var joke Joke
	urlValues := url.Values{}
	urlValues.Set("firstName",uinames.Name)
	urlValues.Set("lastName",uinames.Surname)
	urlValues.Set("limiTo","[nerdy]")
	queryurl := urlValues.Encode()
	joke_uri := "http://api.icndb.com/jokes/random?"+queryurl

	//log.Println("Joke URI:",joke_uri)
	res2, err2 := http.Get(joke_uri)
	if err2 != nil {
		Error(&w, 408, "Chuck Norris has taken down the internet Unable to Fetch Jokes", "Unable to Fetch Jokes", nil)
		
	} else {
		err2 = json.NewDecoder(res2.Body).Decode( &joke)
		if err2 != nil {
			Error(&w, 415, "cannot decode Chuck Norris has used CN2048 encryption", "MalFormed response", nil)

		} else {
			Success(&w, joke.Val.Joke)
		}
	}
}

/*
 * Handles Success reponse to the Http Request
 * 
*/

func Success(w *http.ResponseWriter, result interface{}) {
	writer := *w

	mresult, err := json.Marshal(result)

	if err != nil {
		Error(w, 500, "Internal Server Error", "Error marshalling response JSON", err)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write(mresult)
}


/*
 * Handles Error reponse to the Http Request
 * 
*/

func Error(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
	errorMessage := ""
	writer := *w

	if err != nil {
		errorMessage = err.Error()
	}

	log.Println(logMessage, errorMessage)
	writer.WriteHeader(code)
	writer.Write([]byte(responseText))
}

