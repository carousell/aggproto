## AggProto

<p>
A code generator that enables exposing custom client facing gRPC APIs (Protobuf APIs + server side code)
by consuming registered upstream services as defined by their Protobuf specifications. 
</p>

<p>
The intent of this library is to generate APIs and server code that:
</p>
<b>Given:</b>
<ul>
<li>A repository of internal APIs described in Protobuf format</li>
<li>API Specs defining the behavior of the desired Aggregate endpoints</li>
</ul>
<b>Generates:</b>
<ul>
<li>Protobuf APIs that describe the generated endpoint and that can be used by clients to invoke the endpoints</li>
<li>Server side code that implements the said endpoint and internally invokes the internal apis in the order determined</li>
</ul>

### General workflow
##### Step 1: Register upstream Protobuf
<p>Registration of upstream protos, is required to identify the various message definitions, the sub-definitions and the operations defined by the proto.
This step requires an empty output directory that will be used in later steps to create custom APIs</p>
<p>Usage:</p>

```
protoc --registerproto_out=registry_path=<PATH_TO_REGISTRY>:. -I=<PATH_TO_PROTO_DIR>  <PATH_TO_PROTO_FILE>
```
<p>For example, when registering one of the example proto files in this directory from the root of this project, run: </p>

```
protoc --registerproto_out=registry_path=examples/registry:. -I=examples/protos examples/protos/listing/listing.proto
```

##### Step 2: Define a new API spec
An API Spec is defined as follows:
```yaml
api:
  name: <name of the custom api>
  group: <a package or a team that this api is grouped with>
  version: <an int representing the current version of the spec>
meta:
  goPackage: <the go package name that you want all imports to use>
input:
  - <each line defines either an input/ an alias for the input/ or an output redirection>
output:
  - <each line defines either an output selection or an aliasing on the output>
operations:
  - <you can optionally whitelist operations>
```
For a more detailed description of the api spec look at the details below

##### Step 3: Sync the spec directory
This step processes all the specs found in a directory and generates the client facing protobuf files and the server side go module implementing the generated gRPC service.
<p>Usage:</p>

```
aggproto sync --api_specs_path <PATH_TO_SPECS> --registry_path <PATH_TO_REGISTRY> --go_out_path <PATH_TO_GENERATED_GO> --proto_out_path <PATH_TO_GENERATED_PROTO>
```

<p>For example, when generating the example specs, run the following from the root of this repo:</p>

```
aggproto sync --api_specs_path examples/specs --registry_path examples/registry --go_out_path examples/goOut --proto_out_path examples/protoOut
```

### Examples

Check the examples directory for a detailed list of examples. In order to build the examples, run
`make examples` from the root of the repo.
##### Basic example
The most basic example involves creating a static endpoint that has static values and takes nothing as input.
When we add `listing.title="iPhone"`, this creates a message called listing with a string field called title and the server
code will always return the value "iPhone".
```yaml
api:
  name: mock_listing
  group: static_primitives
  version: 1
meta:
  goPackage: github.com/carousell/aggproto/examples/goOut
output:
  - listing.title="iPhone"
  - listing.description="BNIB iPhone X"
```

##### A more complex example
<p>This is a more real example where we take an id of a listing from the request, use that to fetch a listing. From the 
listing we then extract a `media_id` which is used to fetch a media object. Finally the listing response and media response are 
merged to create the output response desired. </p>

<b>Alias operation</b>
<br>
The `=` operator is referred to in the documentation as an aliasing operation. In this example,
we specify `get_listing.id=listing.GetListingRequest.listing_id` . This indicates that we expect one of the operations inferred 
from the output field selection to require `listing.GetListingRequest.listing_id` as an input parameter and that we are creating an alias for this 
field to instead read `get_listing.id` in the client facing proto. This is valid in both input and output field specifications.
<br>

<b>Pipe operation</b>
<br>
The `<-` operator is referred to as the pipe operation, wherein the output of one operation is used to form the input of another operation.
In this case, we pass the media_id fetched from the listing response to the request of the get media call. This is valid in the input specifications only.


```yaml
api:
  name: masked_listing_w_media
  group: inferred_input
  version: 1
meta:
  goPackage: github.com/carousell/aggproto/examples/goOut
input:
  - get_listing.id=listing.GetListingRequest.listing_id
  - media.GetMediaRequest.media_id<-listing.GetListingResponse.listing.media_id
output:
  - listing.title=listing.GetListingResponse.listing.title
  - listing.description=listing.GetListingResponse.listing.description
  - listing.photo_url=media.GetMediaResponse.media.photo_url
```


## About

<a href="https://github.com/carousell/" target="_blank"><img src="https://avatars2.githubusercontent.com/u/3833591" width="100px" alt="Carousell" align="right"/></a>

**AggProto** is created and maintained by [Carousell](https://carousell.com/). Help us improve this project! We'd love the [feedback](https://github.com/carousell/aggproto/issues) from you.

We're hiring! Find out more at <http://careers.carousell.com/>

## License

**AggProto** is released under Apache License 2.0.
See [LICENSE](https://github.com/carousell/aggproto/blob/master/LICENSE) for more details.

