## To configure SendGrid webhook to receive status callbacks, follow these steps:

```
 Create a webhook endpoint: Set up a server or an endpoint on your server to receive the webhook callbacks from SendGrid. Make sure the endpoint is publicly accessible and can handle incoming HTTP requests.
```
## Configure the webhook settings in SendGrid:
```
1> Log in to your SendGrid account.
```
```
2> Go to the SendGrid dashboard.
```
```
3> Click on "Settings" in the left sidebar.
```
![image](https://github.com/gic-anurag/sendgrid-webhook/assets/89963141/98d9b3f0-062e-4423-9242-1cf1742efdf4)

```
4> Select "Mail Settings" from the dropdown menu.
```
![image](https://github.com/gic-anurag/sendgrid-webhook/assets/89963141/92c835b9-442c-4073-b2b7-9be45e73a500)

```
5> Go to webhook stting's there is event webhooks.
```
![image](https://github.com/gic-anurag/sendgrid-webhook/assets/89963141/5636612b-e30a-4e5b-98bd-4d196bd0fc3f)

```
6> Now create new webhook
```
![image](https://github.com/gic-anurag/sendgrid-webhook/assets/89963141/fc461165-360e-4755-b59e-5ffcbdcf585e)

```
7> In the "HTTP POST URL" field, enter the URL of your webhook endpoint.
```
```
8>Choose the types of events you want to receive notifications for (e.g., delivered, opened, clicked, bounced, etc.).
```
```
9> Click the "Save" button to apply the changes.
```
![image](https://github.com/gic-anurag/sendgrid-webhook/assets/89963141/9197bfbc-2cd3-47ea-8932-9baf28a0b9bc)


## NOTE:-
```
Handle incoming webhook events: Implement the logic on your server to handle incoming webhook events from SendGrid. Parse the JSON payload sent by SendGrid and extract the relevant information such as event type, recipient, message ID, etc. You can use the previously provided code snippet as a starting point and modify it according to your requirements.

Verify event authenticity: To ensure the authenticity of the webhook events, SendGrid provides a signature header (X-Twilio-Email-Event-Webhook-Signature) that you can validate. Implement the verification logic using the provided signature to ensure that the events are indeed sent by SendGrid.

Process the webhook events: Once you receive the webhook events and validate their authenticity, you can process them according to your application's needs. For example, you can update your internal database with the delivery status, send notifications, or perform other actions based on the event type.

By following these steps, you will be able to configure SendGrid to send webhook callbacks to your specified endpoint, allowing you to receive and process status updates for your sent emails.
```

