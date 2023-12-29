import React from 'react';
import { Link } from 'react-router-dom';

export const TopPage: React.FC = () => {
  return (
    <div className="App">
      <Link to="/users">ユーザー一覧</Link>
      <br />
      <Link to="/user/1">個別ユーザー</Link>
      <br />
      <Link to="/user/register">ユーザー登録</Link>
      <br />
    </div>
  );
};
