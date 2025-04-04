import { useMutation, useQueryClient } from "@tanstack/react-query";

import { ACCESS_TOKEN, REFRESH_TOKEN } from "~/lib/constants";
import { ky } from "~/lib/ky-with-auth";
import type { LoginFormSchema } from "~/schemas/login";
import type { LoginResponse } from "~/types/auth-response";

export function useLoginMutation() {
	const queryClient = useQueryClient();

	return useMutation({
		mutationKey: ["login"],
		mutationFn: async ({ email, password }: LoginFormSchema) => {
			return ky
				.post("login", {
					json: {
						email,
						password,
					},
				})
				.json<LoginResponse>();
		},
		onSuccess: (data) => {
			queryClient.setQueryData(["auth"], { user: data.data.user });
			sessionStorage.setItem(ACCESS_TOKEN, data.data.token.access_token);
			sessionStorage.setItem(REFRESH_TOKEN, data.data.token.refresh_token);
		},
		onError: () => {
			sessionStorage.removeItem(ACCESS_TOKEN);
			sessionStorage.removeItem(REFRESH_TOKEN);
		},
	});
}
