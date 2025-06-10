import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import path from 'path';

const PROTO_PATH = path.join(__dirname,  'store.proto');

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

const proto = grpc.loadPackageDefinition(packageDefinition) as any;

const client = new proto.storepb.GeoStore(
  'localhost:50051',
  grpc.credentials.createInsecure()
);

export const storeData = (request: any, p0: (err: any, response: any) => void): Promise<any> => {
  return new Promise((resolve, reject) => {
    client.StoreData(request, (err: any, response: any) => {
      if (err) return reject(err);
      resolve(response);
    });
  });
};

export const getData = (deviceId: string): Promise<any> => {
  return new Promise((resolve, reject) => {
    client.GetData({ deviceId }, (err: any, response: any) => {
      if (err) return reject(err);
      resolve(response);
    });
  });
};

export const findNearby = (
  latitude: number,
  longitude: number,
  radius: number
): Promise<any> => {
  const request = { latitude, longitude, radius };

  return new Promise((resolve, reject) => {
    client.FindNearby(request, (err: any, response: any) => {
      if (err) return reject(err);
      resolve(response);
    });
  });
};

export const healthCheck = (): Promise<any> => {
  return new Promise((resolve, reject) => {
    client.HealthCheck({}, (err: any, response: any) => {
      if (err) return reject(err);
      resolve(response);
    });
  });
};