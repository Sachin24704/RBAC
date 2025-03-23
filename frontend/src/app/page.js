import UserCard from "@/components/UserCard";

export default function Home() {
  const users = [
    {
      id: 1,
      name: "ram",
    },
    {
      id: 2,
      name: "sham",
    },
  ];
  return (
    <div className="flex-col mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">RBAC system</h1>
      {users.map((user) => (
        <UserCard key={user.id} user={user} />
      ))}
    </div>
  );
}
