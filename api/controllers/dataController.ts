import { Request, Response, RequestHandler } from "express"
import { IncomingData, PointsNear } from "../types/data"
import { storeData, getData, findNearby } from "../grpc/client"
import { convertToGrpcStoreRequest, convertToGrpcGetRequest, convertToGrpcNearbyRequest } from "../utils/convertor"

export const postData: RequestHandler = async (_req: Request, _res: Response) => {
    try {
        const grpcRequest = convertToGrpcStoreRequest(_req.body);
        storeData(grpcRequest, (err, response) => {
            if (err) {
                _res.send(500).send({ error: err.message });
            } else {
                _res.send({ status: response.status });
            }
        });
    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}


export const checkRecent: RequestHandler = async (_req: Request, _res: Response) => {
    try {
        const grpcRequest = convertToGrpcGetRequest(_req.body); 
        const request = grpcRequest.deviceId;
        const response = await getData(request); 
        _res.send({
            point : response.point,
            status : response.status
        })

    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}


export const nearbyPoints: RequestHandler = async (_req: Request, _res: Response) => {
    try {
        const grpcRequest = convertToGrpcNearbyRequest(_req.body);
        const response = await findNearby(
            grpcRequest.latitude,
            grpcRequest.longitude,
            grpcRequest.radius
        );

        _res.send({
            status: response.status,
            count: response.count,
            points: response.points, 
        });

    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}

