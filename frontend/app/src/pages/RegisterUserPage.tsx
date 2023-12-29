import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

interface UserData {
  Name: string;
  Email: string;
  Password: string;
}

export const RegisterUserPage = () => {
  const [formData, setFormData] = useState<UserData>({ Name: '', Email: '', Password: '' });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    axios
      .post('http://localhost:8080/users', formData)
      .then((response) => {
        console.log('User registered', response);
        navigate('/users/');
      })
      .catch((error) => {
        console.error('Error registering user: ', error);
        alert('ユーザー登録失敗！');
      });
  };

  return (
    <div>
      <h1>ユーザー登録</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">名前:</label>
          <input type="text" id="name" name="Name" value={formData.Name} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="email">メール:</label>
          <input type="email" id="email" name="Email" value={formData.Email} onChange={handleChange} />
        </div>
        <div>
          <label htmlFor="password">パスワード:</label>
          <input type="password" id="password" name="Password" value={formData.Password} onChange={handleChange} />
        </div>
        <button type="submit">登録</button>
      </form>
    </div>
  );
};
