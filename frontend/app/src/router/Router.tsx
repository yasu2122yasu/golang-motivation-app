import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { IndexUserPage } from '../pages/IndexUserPage';
import { DetailUserPage } from '../pages/DetailUserPage';
import { RegisterUserPage } from '../pages/RegisterUserPage';
import { TopPage } from '../pages/TopPage';

export const Router = () => {
  return (
    <Routes>
      <Route path="/" element={<TopPage />} />
      <Route path="/users" element={<IndexUserPage />} />
      <Route path="/user/1" element={<DetailUserPage />} />
      <Route path="/user/register" element={<RegisterUserPage />} />
    </Routes>
  );
};
