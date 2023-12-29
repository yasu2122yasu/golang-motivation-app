import React, { useState, useEffect } from 'react';

import axios from 'axios';

interface UserData {
  Id: number;
  Name: string;
  Email: string;
  Password: string;
}

export const IndexUserPage = () => {
  const [userData, setUserData] = useState<UserData[]>([]);

  useEffect(() => {
    axios
      .get('http://localhost:8080/users')
      .then((response) => {
        setUserData(response.data); // レスポンスデータを配列として設定
      })
      .catch((error) => {
        console.error('Error fetching data: ', error);
      });
  }, []);

  if (userData.length === 0) return <div>Loading...</div>;

  return (
    <div>
      <h1>ユーザー一覧</h1>
      {userData.map((user) => (
        <div key={user.Id}>
          <p>Id: {user.Id}</p>
          <p>Name: {user.Email}</p>
          <p>Email: {user.Email}</p>
          <p>Password: {user.Password}</p>
        </div>
      ))}
    </div>
  );
};
