import { toast } from "react-toastify";
import { FieldValues } from "react-hook-form/dist/types";
import { createAsyncThunk, isAnyOf } from "@reduxjs/toolkit";
import { createSlice } from "@reduxjs/toolkit";
import { User } from "./../../app/models/user";
import agent from "../../app/api/agent";
import { router } from "../../app/router/Routes";
import jwt_decode from "jwt-decode";
interface AccountState {
  user: User | null;
}

const initialState: AccountState = {
  user: localStorage.getItem("user")
  ? JSON.parse(localStorage.getItem("user")!)
  : null,
};

export const signInUser = createAsyncThunk<User, FieldValues>(
  "account/signInUser",
  async (data, thunkAPI) => {
    try {
      let token = await agent.Account.login(data);
      token = jwt_decode(token);
      let role = token["role"];
      localStorage.setItem(
        "user",
        JSON.stringify({
          email: token["email"],
          userID: token["userID"],
          token: token,
          roles: role === 1 ? ["USER"] : ["ADMIN"],
        })
      );
      return {
        email: token["email"],
        token: token,
        roles: role === 1 ? ["USER"] : ["ADMIN"],
      };
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  }
);

export const fetchCurrentUser = createAsyncThunk<User>(
  "account/fetchCurrentUser",
  async (_, thunkAPI) => {
    thunkAPI.dispatch(setUser(JSON.parse(localStorage.getItem("user")!)));
    try {
      const userDto = await agent.Account.currentUser();
      const { ...user } = userDto;
      return user;
    } catch (error: any) {
      return thunkAPI.rejectWithValue({ error: error.data });
    }
  },
  {
    condition: () => {
      if (!localStorage.getItem("user")) return false;
    },
  }
);

export const accountSlice = createSlice({
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
    builder.addCase(fetchCurrentUser.rejected, (state) => {
      state.user = null;
      localStorage.removeItem("user");
      toast.error("Session expired - please login again");
      router.navigate("/");
    });
    builder.addMatcher(
      isAnyOf(signInUser.fulfilled, fetchCurrentUser.fulfilled),
      (state, action) => {
        state.user = {
          ...action.payload,
        };
      }
    );
    builder.addMatcher(
      isAnyOf(signInUser.rejected, fetchCurrentUser.rejected),
      (state, action) => {
        throw action.payload;
      }
    );
  },
});

export const { signOut, setUser } = accountSlice.actions;
