import { Request, Response, RequestHandler } from "express"
import { IncomingData, PointsNear } from "../types/data"

export const postData: RequestHandler = async (_req: Request, _res: Response) => {
    try {

    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}


export const checkRecent: RequestHandler = async (_req: Request, _res: Response) => {
    try {

    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}


export const nearbyPoints: RequestHandler = async (_req: Request, _res: Response) => {
    try {

    } catch (err) {
        _res.status(500).json({ message: "Internal Server Error occured" })
    }
}

