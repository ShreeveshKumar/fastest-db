import { Router } from "express";
import getHealth from "../controllers/healthController";
import { postData, checkRecent, nearbyPoints } from "../controllers/dataController"
import { validateInput } from "../middlewares/validateInput";

const router = Router();


router.get("/heath", getHealth);
router.post("/data", validateInput, postData);
router.get("/data:id", checkRecent);
router.get("/nearby", nearbyPoints);


export default router; 