package main

import "net/http"

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	feed, err := app.store.Posts.GetUserFeed(ctx, 100)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusNoContent, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}
