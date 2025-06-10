
import express, { Router } from "express";
import * as http from "http";
import dotenv from "dotenv";
import helmet from "helmet";
import dataRoutes from "./routes/index"

dotenv.config({
    path: "./.env"
});

const app = express();
const PORT = process.env.PORT || 8000;
app.use(express.json());
app.use(helmet());



app.use("/api/v1", dataRoutes);

const server = http.createServer(app);

server.listen(PORT, () => {
    console.log(`Server is running on ${PORT}`);
})