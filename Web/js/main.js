const BASE_URL = "https://iiaeky1f01.execute-api.us-east-1.amazonaws.com/prod/";
const TRIGGER_URL = "lambdatrigger";

function toGW() {
	fetch(BASE_URL + TRIGGER_URL, {
		method: 'POST',
		body: '{"name":"Bimesh"}',
		headers:{
			'Content-Type': 'application/json'
		}
	})
	.then(data => data.json())
	.then(json => console.log(json))
	.catch(error => console.log(error));
}