"use client";

import { useState, useEffect } from "react";
import RoleCard from "./RoleCard";

const UserCard = ({ user }) => {
  const [roles, setRoles] = useState([]);
  useEffect(() => {
    const fetchPermissions = async () => {
      try {
        const response = await fetch(
          `http://localhost:8000/users/${user.id}/permissions`
        );
        const data = await response.json();
        setRoles(data.roles);
      } catch (error) {
        console.error("error fetching permissions", error);
      }
    };
    fetchPermissions();
  }, [user.id]);
  return (
    <div className="p-4 mb-4 rounded-2xl shadow">
      <h2 className="text-xl font- bold mb-2">{user.name}</h2>
      {roles.map((role) => (
        <RoleCard role={role} key={role.id} />
      ))}
    </div>
  );
};

export default UserCard;
