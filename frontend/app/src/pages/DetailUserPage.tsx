import React, { useState, useEffect } from 'react';
import axios from 'axios';

interface UserData {
  Id: number;
  Name: string;
  Email: string;
  Password: string;
}

export const DetailUserPage = () => {
  const [userData, setUserData] = useState<UserData | null>(null);

  useEffect(() => {
    axios
      .get('http://localhost:8080/users/1')
      .then((response) => {
        setUserData(response.data);
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, []);

  if (!userData) return <div>Loading...</div>;

  return (
    <div>
      <h1>個別ユーザー情報</h1>
      <div>
        <p>Id: {userData.Id}</p>
        <p>Name: {userData.Name}</p>
        <p>Email: {userData.Email}</p>
        <p>Password: {userData.Password}</p>
      </div>
    </div>
  );
};
