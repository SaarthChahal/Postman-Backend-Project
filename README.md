



For user login authentication, I am planning to integrate user verification through their Google account. 

To create login ID, users can either give their phone numbers, or their e mail addresses.
The next step is to verify the given information. 

Sign in with Google is based on OAuth 2.0. Carry out the following steps:

1. Setting up OAuth 2.0
Set up a project in the Google API Console to obtain OAuth 2.0 credentials, set a redirect URI, and (optionally) customize the branding information that  users see on the user-consent screen. 

2.Obtain OAuth 2.0 credentials
You need OAuth 2.0 credentials, including a client ID and client secret, to authenticate users and gain access to Google's APIs.

3.Set a redirect URI
The redirect URI that you set in the API Console determines where Google sends responses to your authentication requests.

4. Customize the user consent screen

5. Authenticating the user
Authenticating the user involves obtaining an ID token and validating it.

Server flow
Make sure you set up your app in the API Console to enable it to use these protocols and authenticate your users. When a user tries to log in with Google, you need to:
1.	Create an anti-forgery state token
2.	Send an authentication request to Google
3.	Confirm the anti-forgery state token
4.	Exchange code for access token and ID token
5.	Obtain user information from the ID token
6.	Authenticate the user




