import { createBrowserRouter, Navigate } from "react-router-dom";
import FlightsPage from "../../features/flightsPage/FlightsPage";
import Login from "../../features/account/Login";
import Register from "../../features/account/Register";

import App from "../layout/App";
import RequireAuth from "./RequireAuth";

export const router = createBrowserRouter([
    {
        path: '/',
        element: <App />,
        children: [
            // authenticated routes za loginovanog usera
            {element: <RequireAuth />, children: [
            
            ]},
            // admin routes
            {element: <RequireAuth roles={[]} />, children: [
                    
            ]},
            {path: 'flights', element: <FlightsPage />},  
            {path: 'login', element: <Login />},
            {path: 'register', element: <Register />},
            {path: '*', element: <Navigate replace to='/not-found' />}
        ]
    }
])