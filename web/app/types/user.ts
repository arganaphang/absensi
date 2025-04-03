type UserRole = "admin" | "staff";

export type User = {
  id: string;
  email: string;
  fullname: string;
  birthdate: Date;
  position: string;
  phone: string;
  address: string;
  role: UserRole;
};
