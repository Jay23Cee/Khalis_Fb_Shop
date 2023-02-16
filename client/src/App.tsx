import React from 'react';

import {useQuery, gql } from '@apollo/client';



// Example of a GraphQL query
const query = gql`
  query {
  facebook {
    id
    posts {
      data {
        id
        fullPicture
        comments {
          data {
            id
            createdTime
            message
            from {
              id
              name
            }
            comments{
              data{
                 id
            createdTime
            message
            from {
              id
              name
            }
              }
            }
          }
          paging {
            cursors {
              before
              after
            }
          }
        }
      }
      paging {
        cursors {
          before
          after
        }
      }
    }
  }
}

`;


interface FacebookData {
  id: string;
  posts: {
    data: {
      id: string;
      fullPicture?: string;
      full_picture?: string;
      comments?: {
        data: CommentData[];
        paging?: {
          cursors?: {
            before: string;
            after: string;
          };
        };
      };
      paging?: {
        cursors?: {
          before: string;
          after: string;
        };
      };
    }[];
    paging?: {
      cursors?: {
        before: string;
        after: string;
      };
    };
  };
}

interface CommentData {
  id: string;
  created_time: string;
  message?: string;
  from?: {
    id: string;
    name: string;
  };
  attachment?: {
    media: {
      image: {
        height: number;
        src: string;
        width: number;
      };
    };
    target: {
      id: string;
      url: string;
    };
    type: string;
    url: string;
  };
  comments?: {
    data: CommentData[];
    paging?: {
      cursors?: {
        before: string;
        after: string;
      };
    };
  };
}




function Comment({ comment }: { comment: CommentData }) {
  return (
    <li key={comment.id}>
      <p>
        {comment.from?.name}: {comment.message}
      </p>
      {comment.comments?.data && (
        <ul>
          {comment.comments.data.map((nestedComment) => (
            <Comment key={nestedComment.id} comment={nestedComment} />
          ))}
        </ul>
      )}
    </li>
  );
}

function Post({ post }: { post: FacebookData['posts']['data'][0] }) {
  
  console.log(post.fullPicture, "POST FULL PICTURE HERE")
  return (
    <li key={post.id}>
      <img src={post.fullPicture} alt="Post image" />
      <p>{post.full_picture} !POST FULL PICTURE</p>
      {post.comments?.data && (
        <ul>
          {post.comments.data.map((comment: any) => (
            <Comment key={comment.id} comment={comment} />
          ))}
        </ul>
      )}
    </li>
  );
}

function App() {
  const { loading, error, data } = useQuery<{ facebook: FacebookData }>(query);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  // Add a check to ensure that `data` is defined before accessing its properties
  if (!data || !data.facebook || !data.facebook.posts || !data.facebook.posts.data) {
    return null;
  }
  console.log('data', data);
  return (
    <div>
      <h1>My App</h1>
      <ul>
        {data.facebook.posts.data.map((post) => (
          <Post key={post.id} post={post} />
        ))}
      </ul>
    </div>
  );
}


export default App;
