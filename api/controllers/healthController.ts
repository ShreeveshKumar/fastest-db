import { Request, Response, RequestHandler } from "express";
import { Healthstatus } from "../types/data";

const getHealth: RequestHandler = async (_req: Request, _res: Response) => {
    try {
        _res.status(200).json({ status: "healthy" });
    } catch (err) {
        _res.status(500).json({ message: "Internal server error occured" });
    }
}

export default getHealth; 

