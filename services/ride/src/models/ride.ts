import * as yup from "yup";

export const rideSchema = yup.object({
  city: yup.string().required(),
  distance: yup
    .number()
    .positive()
    .required(),
  time: yup
    .number()
    .positive()
    .required()
});

export type Ride = yup.InferType<typeof rideSchema>;
