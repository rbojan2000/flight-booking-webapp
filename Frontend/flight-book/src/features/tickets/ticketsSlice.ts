import { toast } from "react-toastify";
import { FieldValues } from "react-hook-form/dist/types";
import { createAsyncThunk, isAnyOf } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import { User } from "./../../app/models/user";
import agent from "../../app/api/agent";
import { router } from "../../app/router/Routes";

interface AccountState {
  user: User | null;
}

const initialState: AccountState = {
  user: null,
};

export const signInUser = createAsyncThunk<User, FieldValues>(
  "account/signInUser",
  async (data, thunkAPI) => {
    try {
        window.alert('dsdsds')
      const userDto = await agent.Account.login(data);
      console.log(userDto);
      const { ...user } = userDto;
      return user;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);


export const ticketSlice = createSlice({
  name: "account",
  initialState,
  reducers: {
    signOut: (state) => {
      state.user = null;
      localStorage.removeItem("user");
      router.navigate("/");
    },
    setUser: (state, action) => {
      state.user = action.payload;
      let claims = JSON.parse(atob(action.payload.token.split(".")[1]));
      let roles =
        claims["http://schemas.microsoft.com/ws/2008/06/identity/claims/role"];
      state.user = {
        ...action.payload,
        roles: typeof roles === "string" ? [roles] : roles,
      };
    },
  },
  extraReducers: (builder) => {
        
    
    builder.addMatcher(
      isAnyOf(),
      (state, action) => {
       
        }

    );
  },
});

export const { signOut, setUser } = ticketSlice.actions;
