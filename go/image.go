package main

// func getImage(w http.ResponseWriter, r *http.Request) {
// 	pidStr := r.PathValue("id")
// 	pid, err := strconv.Atoi(pidStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	post := Post{}
// 	err = db.Get(&post, "SELECT * FROM `posts` WHERE `id` = ?", pid)
// 	if err != nil {
// 		log.Print(err)
// 		return
// 	}

// 	ext := r.PathValue("ext")

// 	if ext == "jpg" && post.Mime == "image/jpeg" ||
// 		ext == "png" && post.Mime == "image/png" ||
// 		ext == "gif" && post.Mime == "image/gif" {
// 		w.Header().Set("Content-Type", post.Mime)
// 		_, err := w.Write(post.Imgdata)
// 		if err != nil {
// 			log.Print(err)
// 			return
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusNotFound)
// }
