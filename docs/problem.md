## Problem
1. User should be able to create a CFS with the following information: event number, event
type (with type code), event time, dispatch time, responder.
2. User should be able to search for CFS within a time range with paging and sorting order.
3. User should be able to search for CFS that assigned to a responder.
4. CFS belongs to different agencies are not allowed to be exposed to other agencies.
5. User and responder should belong to only one agency.

```JSON
{
	"agency_id": "4f9b99eb-490a-484e-bade-15e3841dfda9",
	"event_id": "562c89de-f140-4482-8ef5-5f1703b286b6",
	"event_number": "3234019",
	"event_type_code": "SMO",
	"event_time": "2020-11-25 07:36:04.193",
	"dispatch_time": "2020-11-26 13:55:46.466",
	"responder": "OFFICER_001"
}
```

Goal:
 - Your task is to design CFS service to support the above requirements and implement the API to
support CFS search by time range.

## Requirements
As a Senior Engineer, you are going to take lead a team of talented people, leading the designs
and implements phase that involve multiple features or components to ensure the efficient end-
to-end, integration and component of our product.
- You are going to create a well-written technical design document includes
	o System design, with detailed reasons to support your decision
	o Data models and database design.
	o API contracts
- You are going to define a test plan, define metrics that measure quality of this products,
break it down, and come up with a plan to support your team deliver it

## Expectation
Please prepare a document to describe your plan to tackle on every requirements above.

• We’d like to have an understanding about your technical aptitude, so we highly prefer
you provide a Github project, where you define your selected tools, as well as how you’re
going to setup your project to kick off the product
• We’d like to see your plan to delivery this product