
This application for operators in order to send a notification to the customer about the purchase. 

Notifications can go through via mail and via SMS.

This will depend on the choice of the operator. 

All errors that occur during sending will be shown to the operator, as well as save a storage database.

Customer has: id, phone_number and email.
Purchase has: id, product and price.

In order to build this application:

 go build cmd/main.go 

You can run it by your option. option can sms or mail. 

If you type sms it will send sms to customer. 
If you type mail it will send mail to customer.

 ./main option

 Notice: there all phone, mail, password and other informations are written as a example. You should write your owns ones.
