package org.horaapps.leafpic.data;

import android.content.Context;
import android.media.MediaScannerConnection;
import android.net.Uri;
import android.provider.MediaStore;

import org.horaapps.leafpic.util.StringUtils;
import org.horaapps.leafpic.util.file.DeleteException;

import java.io.File;
import java.util.ArrayList;

import io.reactivex.Observable;

/**
 * Created by dnld on 8/8/17.
 */

public class MediaHelper {
    private static Uri external = MediaStore.Files.getContentUri("external");



    public static Observable<Media> deleteMedia(Context context, Media mediaToDelete) {
        return Observable.create(subscriber -> {
            boolean deleteSuccess = internalDeleteMedia(context, mediaToDelete);
            if (deleteSuccess) subscriber.onNext(mediaToDelete);
            else subscriber.onError(new DeleteException(mediaToDelete));
            subscriber.onComplete();
        });
    }


    public static Observable<Media> deleteMedia(Context context, ArrayList<Media> mediaToDelete) {
        return Observable.create(subscriber -> {
            for (Media media : mediaToDelete) {
                boolean deleteSuccess = internalDeleteMedia(context, media);
                if (deleteSuccess) subscriber.onNext(media);
                else subscriber.onError(new DeleteException(media));
            }
            subscriber.onComplete();
        });
    }

    private static boolean internalDeleteMedia(Context context, Media media) {
        File file = new File(media.getPath());
        boolean success = StorageHelper.deleteFile(context, file);
        if (success)
            context.getContentResolver().delete(external,
                    MediaStore.MediaColumns.DATA + "=?", new String[]{file.getPath()});
        return success;
    }

    public static boolean renameMedia(Context context, Media media, String newName) {
        boolean success = false;
        try {
            File from = new File(media.getPath());
            File to = new File(StringUtils.getPhotoPathRenamed(media.getPath(), newName));
            if (success = StorageHelper.moveFile(context, from, to)) {
                context.getContentResolver().delete(external,
                        MediaStore.MediaColumns.DATA + "=?", new String[]{from.getPath()});

                scanFile(context, new String[]{to.getAbsolutePath()});
                media.setPath(to.getAbsolutePath());
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return success;
    }

    public static boolean moveMedia(Context context, Media media, String targetDir) {
        boolean success = false;
        try {
            File from = new File(media.getPath());
            File to = new File(targetDir, from.getName());
            if (success = StorageHelper.moveFile(context, from, to)) {

                context.getContentResolver().delete(external,
                        MediaStore.MediaColumns.DATA + "=?", new String[]{from.getPath()});


                scanFile(context, new String[]{StringUtils.getPhotoPathMoved(media.getPath(), targetDir)});
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return success;
    }

    public static boolean copyMedia(Context context, Media media, String targetDir) {
        boolean success = false;
        try {
            File from = new File(media.getPath());
            File to = new File(targetDir);
            if (success = StorageHelper.copyFile(context, from, to))
                scanFile(context, new String[]{StringUtils.getPhotoPathMoved(media.getPath(), targetDir)});

        } catch (Exception e) {
            e.printStackTrace();
        }
        return success;
    }

    private static void scanFile(Context context, String[] path) {
        MediaScannerConnection.scanFile(context, path, null, null);
    }
}
