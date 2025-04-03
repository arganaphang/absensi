import { useQueryClient } from "@tanstack/react-query";
import { useEffect } from "react";
import { redirect, useNavigate } from "react-router";

import { authQueryOptions, useAuthQuery } from "~/api/auth-query";
import { useLogoutMutation } from "~/api/logout-mutation";
import type { AuthResponse } from "~/types/auth-response";

type AuthState =
  | { user: null; status: "PENDING" }
  | { user: null; status: "UNAUTHENTICATED" }
  | { user: AuthResponse["data"]["user"]; status: "AUTHENTICATED" };

type AuthUtils = {
  signIn: () => void;
  signOut: () => void;
  ensureData: () => Promise<AuthResponse | undefined>;
};

type AuthData = AuthState & AuthUtils;

function useAuth(): AuthData {
  const authQuery = useAuthQuery();
  const signOutMutation = useLogoutMutation();
  let navigate = useNavigate();

  const queryClient = useQueryClient();

  useEffect(() => {
    console.log("authQuery.data", authQuery.data);
    navigate(".", { replace: true });
  }, [authQuery.data]);

  useEffect(() => {
    if (authQuery.error === null) return;
    queryClient.setQueryData(["auth"], null);
  }, [authQuery.error, queryClient]);

  const utils: AuthUtils = {
    signIn: () => {
      redirect("/login");
    },
    signOut: () => {
      signOutMutation.mutate();
    },
    ensureData: () => {
      return queryClient.ensureQueryData(authQueryOptions());
    },
  };

  switch (true) {
    case authQuery.isPending:
      return { ...utils, user: null, status: "PENDING" };

    case !authQuery.data:
      return { ...utils, user: null, status: "UNAUTHENTICATED" };

    default:
      return {
        ...utils,
        user: authQuery.data.data.user,
        status: "AUTHENTICATED",
      };
  }
}

export { useAuth };
export type { AuthData };
