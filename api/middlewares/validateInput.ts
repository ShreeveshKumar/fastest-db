import { Response, Request, RequestHandler, NextFunction } from "express";
import { IncomingData } from "../types/data";


export const validateInput: RequestHandler = async (_req: Request, _res: Response, next: NextFunction) => {
    try {
        const { latitude, longitude, timeStamp, deviceId } = _req.body as IncomingData;

        if (!latitude || !longitude || !timeStamp || !deviceId) {
            _res.status(400).send({ error: 'Invalid input format. Expected { x, y, timestamp, deviceId }.' });
        }

        next();
    } catch (err) {
        _res.status(500).json({ message: "Internal Server occured" });
    }
}