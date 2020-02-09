import grpc from "grpc";
import messages from "../../../protos/pricing_pb";
import services from "../../../protos/pricing_grpc_pb";

Promise.resolve().then(() => {
  const client = new services.PricingClient(
    "pricing-api:50051",
    grpc.credentials.createInsecure()
  );

  const request = new messages.GetPricingFeesByCityRequest();
  request.setCity("SAO_PAULO");

  client.getPricingFeesByCity(request, (err, res) => {
    console.log("Response from Golang: ", res?.toObject());
  });
});
