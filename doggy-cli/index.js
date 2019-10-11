'use strict'

const axios = require('axios');

exports.getDog = (endpoint, id) => {
  const url = endpoint.url;
  const port = endpoint.port;

  return axios.request({
    method: 'GET',
    baseURL: `${url}:${port}`,
    url: `/dogs/${id}`,
    headers: { Accept: 'application/json;charset=UTF-8' }
  });
};

exports.saveDog = (endpoint, dog) => {
  const url = endpoint.url;
  const port = endpoint.port

  return axios.request({
    method: 'POST',
    baseURL: `${url}:${port}`,
    url: '/dogs',
    headers: {
      'Content-Type': 'application/json;charset=UTF-8',
      Accept: 'application/json;charset=UTF-8'
    },
    data: dog
  });
};

