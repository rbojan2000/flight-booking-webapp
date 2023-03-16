export interface Ticket {
  
    description: string;
}

export interface TicketParams {
    orderBy: string;
    searchTerm?: string;
    types: string[];
    brands: string[];
    pageNumber: number;
    pageSize: number;
}