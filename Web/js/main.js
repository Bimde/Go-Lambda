const BASE_URL = "https://iiaeky1f01.execute-api.us-east-1.amazonaws.com/prod/";
const TRIGGER_URL = "lambdatrigger";

function getValue() {
	return document.getElementById("bimde_text_input").value;
}

function setTextWithValue(s) {
	document.getElementById("bimde_text_field").innerHTML = s;
}

function toGW() {
	var value = getValue();
	fetch(BASE_URL + TRIGGER_URL, {
		method: 'POST',
		body: JSON.stringify({name: value}),
		headers:{
			'Content-Type': 'application/json'
		}
	})
	.then(data => data.json())
	.then(json => {
		console.log(json);
		setTextWithValue("The AWS Lambda function says, \"" + json.response + "\"");
	})
	.catch(error => {
		console.log(error);
		setTextWithValue("Whoops, something went wrong :(<br> Check out the console if you want more details.");
	});
}