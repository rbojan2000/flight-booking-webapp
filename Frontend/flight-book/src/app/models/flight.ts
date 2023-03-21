export interface Flight {
    ID: string;
    Date: string;
    Departure: {
      Country: string;
      City: string;
    };
    Arrival: {
      Country: string;
      City: string;
    };
    PassengerCount: number;
    Capacity: number;
    Price: number;
}