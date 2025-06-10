import { Healthstatus, IncomingData, NearByPoints } from "../types/data";

export const convertToGrpcStoreRequest = (data: IncomingData) => {
  return {
    point: {
      latitude: data.latitude,
      longitude: data.longitude,
      deviceId: data.deviceId,
      timeStamp: data.timeStamp,
    },
  };
};

export const convertToGrpcGetRequest = (deviceId: string) => {
  return {
    deviceId,
  };
};

export const convertToGrpcNearbyRequest = (data: NearByPoints) => {
  return {
    latitude: data.latitude,
    longitude: data.longitude,
    radius: data.radius,
  };
};

export const convertToGrpcHealthCheckRequest = (data: Healthstatus) => {
  return {
    status: data.status, 
  };
};