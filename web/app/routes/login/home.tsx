import { useLoginMutation } from "~/api/login-mutation";
import type { Route } from "./+types/home";
import { Button } from "~/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "~/components/ui/card";
import { Input, InputPassword } from "~/components/ui/input";
import { Label } from "~/components/ui/label";
import { loginFormSchema } from "~/schemas/login";

export function meta({}: Route.MetaArgs) {
  return [{ title: "Login" }];
}

export default function Home() {
  const loginMutation = useLoginMutation();

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const form = e.currentTarget;
    const formData = new FormData(form);
    const data = Object.fromEntries(formData);
    const dataParsed = loginFormSchema.parse(data);
    await loginMutation.mutateAsync(dataParsed);
  }

  return (
    <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
      <div className="w-full max-w-sm">
        <div className="flex flex-col gap-6">
          <Card>
            <CardHeader>
              <CardTitle className="text-2xl">Login</CardTitle>
              <CardDescription>
                Enter your email below to login to your account
              </CardDescription>
            </CardHeader>
            <CardContent>
              <form onSubmit={handleSubmit}>
                <div className="flex flex-col gap-6">
                  <div className="grid gap-2">
                    <Label htmlFor="email">Email</Label>
                    <Input
                      id="email"
                      type="email"
                      placeholder="me@example.com"
                      name="email"
                      required
                    />
                  </div>
                  <div className="grid gap-2">
                    <div className="flex items-center">
                      <Label htmlFor="password">Password</Label>
                    </div>
                    <InputPassword
                      id="password"
                      placeholder="************"
                      name="password"
                      required
                    />
                  </div>
                  <Button type="submit" className="w-full">
                    Login
                  </Button>
                </div>
              </form>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
