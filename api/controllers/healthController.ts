import { Request, Response } from "express";


const getHealth = async (_req: Request, _res: Response) => {
    try {
 

    } catch (err) {
        return _res.status(500).json({ message: "Internal server error occured" });
    }
}



export default getHealth; 

