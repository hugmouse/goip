## goip - serverless alternative to "What is my IP?" services

Based on the Vercel ecosystem. 

It is a service that responds to any request with information about the request, including IP.

## Usage example

```sh
curl https://goip.vercel.app
```

```json
{
	"Method": "GET",
	"URL": {
		"Scheme": "",
		"Opaque": "",
		"User": null,
		"Host": "",
		"Path": "/",
		"RawPath": "",
		"ForceQuery": false,
		"RawQuery": "",
		"Fragment": "",
		"RawFragment": ""
	},
	"Proto": "HTTP/1.1",
	"ProtoMajor": 1,
	"ProtoMinor": 1,
	"Header": {
		"Accept": [
			"*/*"
		],
		"Host": [
			"goip.vercel.app"
		],
		"User-Agent": [
			"curl/7.76.1"
		],
		"X-Forwarded-For": [
			"65.244.137.5"
		],
		"X-Forwarded-Host": [
			"goip.vercel.app"
		],
		"X-Forwarded-Proto": [
			"https"
		],
		"X-Real-Ip": [
			"65.244.137.5"
		],
		"X-Vercel-Deployment-Url": [
			"myip-csb3g0nhy-mysh.vercel.app"
		],
		"X-Vercel-Forwarded-For": [
			"65.244.137.5"
		],
		"X-Vercel-Id": [
			"arn1::crs7s-2147483647-fefefef6c7e5"
		]
	},
	"Body": "",
	"ContentLength": 0,
	"TransferEncoding": null,
	"Host": "goip.vercel.app",
	"Form": null,
	"PostForm": null,
	"MultipartForm": null,
	"Trailer": null,
	"RemoteAddr": "65.244.137.5",
	"RequestURI": "",
	"TLS": null,
	"Response": null
}
```


