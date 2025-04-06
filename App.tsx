import React from 'react';

interface MyComponentProps {
  title: string;
  count: number;
}

const MyComponent: React.FC<MyComponentProps> = ({ title, count }) => {
  return (
    <div>
      <h1>{title}</h1>
      <p>The count is: {count}</p>
    </div>
  );
};

export default MyComponent;
