import { useMutation, useQueryClient } from "@tanstack/react-query";

import { ACCESS_TOKEN, REFRESH_TOKEN } from "~/lib/constants";
import { ky } from "~/lib/ky-with-auth";
import type { LogoutResponse } from "~/types/auth-response";

export function useLogoutMutation() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationKey: ["logout"],
		mutationFn: async () => {
			return ky.post("logout").json<LogoutResponse>();
		},
		onSuccess: () => {
			queryClient.setQueryData(["auth"], null);
			sessionStorage.removeItem(ACCESS_TOKEN);
			sessionStorage.removeItem(REFRESH_TOKEN);
		},
	});
}
