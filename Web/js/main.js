const BASE_URL = "";

function toGW() {
	fetch(BASE_URL)
	.then(data => data.json())
	.then(json => console.log(json))
	.catch(error => console.log(error));
}