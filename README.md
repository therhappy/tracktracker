# TrackTracker
A tracker tool for ETFs

## In short
This project aims to provide a simple client to monitor ETF values and get an overview of an ETF portfolio.

## Structure
The plan ins currently to implement
* a **server** in Go which interacts with a MongoDB database
* an **UI client** adressing the server

An external API is used to retrieve ETF information. The current position on this matter is to use OPCVM360.

## Progress
* A skeleton for the Go sever + mongoDB database has been designed. Right now it can only adress UCs, refered as "Products". *Controller* and *model* are operational, but routes have yet to be defined.
* The *example* file provides a course of action soliciting some of these functions to demonstrate functions purpose. Thorough testing will happen ... someday.
* Requests to the external API are not implemented yet
* Interface reflection has yet to start.