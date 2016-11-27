package main

import "fmt"
// The dumb way to deal with SOAP requests

func userRequest(userId string) string {
   template := soapify(userTemplate)
   return fmt.Sprintf(template, userId)
}

func soapify(template string) string {
   return fmt.Sprintf(soapTemplate, template)
}

const soapTemplate = 
`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ws="http://we&#x2764;namespaces.com/">
   <soapenv:Header/>
   <soapenv:Body>
      %s
   </soapenv:Body>
</soapenv:Envelope>`

const userTemplate = 
`<github:user>
   <arg0>%s</arg0>
</github:user>`

func main() {
  payload := userRequest("dmichael")
	fmt.Println(payload)
}