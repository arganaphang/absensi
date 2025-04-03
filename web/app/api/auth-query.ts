import { queryOptions, useQuery } from "@tanstack/react-query";

import { ky } from "~/lib/ky-with-auth";
import type { AuthResponse } from "~/types/auth-response";

export function authQueryOptions() {
  return queryOptions({
    queryKey: ["auth"],
    queryFn: () => ky.get("profile").json<AuthResponse>(),
  });
}

export function useAuthQuery() {
  return useQuery(authQueryOptions());
}
