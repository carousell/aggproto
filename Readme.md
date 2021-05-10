## AggProto
<p>
Its a common scenario in enterprise software architectures to have a layer thin client facing 
layer which is responsible for the following use cases and more:
</p>
<ul>
<li>Authenticate user requests and check for permissions for access resources</li> 
<li>To invoke various operations each handled by internal services in some order and to generate
an output that may have additional data obtained from other internal services for handling UI concerns</li>
</ul>

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

### Use Cases Handled

<ul>
<li>Masked endpoints: Invoke a single endpoint and get a subset of the fields returned by the internal endpoint</li>
<li>Union endpoints: Invoke two or more independent endpoints and with the request and response composing the union of the independent APIs</li>
<li>Aliased fields: Rename the input and output fields</li>
<li>Transitive input: Pass the output of one operation to the input of another operation in a sequential manner</li>
</ul>


### To Be Handled
<ul>
<li>One of handling</li>
<li>Enum and some other primitive types handling</li>
<li>Repeated fields in input transformations and intermediate operations </li>
<li>A UI to make the spec definition process a bit more intuitive</li>
<li>Additional boilerplate: go.mod + run server code for each group of APIs</li>
<li>Defining and handling concerns like authentication, logging</li>
<li>non field level output def</li>
</ul>