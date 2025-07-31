// src/api/articleAPI.js
import axios from 'axios';

const BASE_URL = 'http://localhost:8080';

export const fetchArticles = async () => {
  try {
    const res = await axios.get(`${BASE_URL}/articles`);
    return res.data.data || [];
  } catch (err) {
    console.error('Failed to fetch articles:', err);
    return [];
  }
};

export const getArticleById = async (id) => {
  try {
    const res = await axios.get(`${BASE_URL}/articles/${id}`);
    return res.data;
  } catch (err) {
    console.error(`Failed to fetch article ${id}:`, err);
    throw err;
  }
};

export const createArticle = async (data) => {
  try {
    const res = await axios.post(`${BASE_URL}/articles`, data);
    return res.data;
  } catch (err) {
    console.error('Failed to create article:', err);
    throw err;
  }
};

export const updateArticle = async (id, data) => {
  try {
    const res = await axios.put(`${BASE_URL}/articles/${id}`, data);
    return res.data;
  } catch (err) {
    console.error(`Failed to update article ${id}:`, err);
    throw err;
  }
};

export const deleteArticle = async (id) => {
  try {
    const res = await axios.delete(`${BASE_URL}/articles/${id}`);
    return res.data;
  } catch (err) {
    console.error(`Failed to delete article ${id}:`, err);
    throw err;
  }
};
