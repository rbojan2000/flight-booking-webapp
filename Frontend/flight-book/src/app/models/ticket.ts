

export interface Ticket {
    ID: string;
    Flight: {
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
    };
    User: {
      ID: string;
      Name: string;
      Surname: string;
      Email: string;
      Password: string;
      Type: number;
      Address: {
        Country: string;
        City: string;
      };
    };
  }