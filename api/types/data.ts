
export interface IncomingData {
    latitude: number;
    longitude: number;
    deviceId: string;
    timeStamp: string;
}


export interface Healthstatus {
    status: Boolean;
}


export interface PointsNear {
    points: Array<any>;
}

export interface NearByPoints{
    latitude: number,
    longitude: number,
    radius: number,
}