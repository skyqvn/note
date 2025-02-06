# FFmpeg 主要功能

FFmpeg 是一个强大的多媒体处理工具，支持视频、音频和图像的编解码、转换、处理等操作。以下是其主要功能：

## 1. **视频/音频转换**
   - 支持多种格式之间的转换，如 MP4、AVI、MKV、MOV、MP3、AAC 等。
   - 示例：
     ```bash
     ffmpeg -i input.mp4 output.avi
     ```

## 2. **视频/音频提取**
   - 从视频中提取音频。
     ```bash
     ffmpeg -i input.mp4 -vn -acodec copy output.mp3
     ```
   - 从视频中提取图像帧。
     ```bash
     ffmpeg -i input.mp4 -vf fps=1 image%d.png
     ```

## 3. **视频/音频裁剪**
   - 根据时间点裁剪视频或音频。
     ```bash
     ffmpeg -i input.mp4 -ss 00:01:00 -to 00:02:00 -c copy output.mp4
     ```

## 4. **视频/音频合并**
   - 将多个视频或音频文件合并为一个文件。
     ```bash
     ffmpeg -i input1.mp4 -i input2.mp4 -filter_complex concat=n=2:v=1:a=1 output.mp4
     ```

## 5. **调整视频分辨率**
   - 调整视频的分辨率。
     ```bash
     ffmpeg -i input.mp4 -vf scale=1280:720 output.mp4
     ```

## 6. **调整视频/音频码率**
   - 调整视频或音频的码率以控制文件大小和质量。
     ```bash
     ffmpeg -i input.mp4 -b:v 1M -b:a 128k output.mp4
     ```

## 7. **添加水印**
   - 在视频中添加图像或文字水印。
     ```bash
     ffmpeg -i input.mp4 -i watermark.png -filter_complex "overlay=10:10" output.mp4
     ```

## 8. **视频/音频压缩**
   - 通过调整编码参数压缩视频或音频文件。
     ```bash
     ffmpeg -i input.mp4 -vcodec libx264 -crf 28 output.mp4
     ```

## 9. **视频/音频格式转换**
   - 将视频或音频转换为其他格式。
     ```bash
     ffmpeg -i input.mp4 -c:v libx265 -c:a aac output.mkv
     ```

## 10. **视频/音频滤镜**
   - 应用各种滤镜效果，如旋转、裁剪、模糊、锐化等。
     ```bash
     ffmpeg -i input.mp4 -vf "rotate=90*PI/180" output.mp4
     ```

## 11. **视频/音频流处理**
   - 提取或删除视频/音频流。
     ```bash
     ffmpeg -i input.mp4 -map 0:v -map 0:a output.mp4
     ```

## 12. **屏幕录制**
   - 录制屏幕并保存为视频文件。
     ```bash
     ffmpeg -f avfoundation -i "1:0" output.mp4
     ```

## 13. **直播推流**
   - 将视频或音频推流到直播平台。
     ```bash
     ffmpeg -re -i input.mp4 -c:v libx264 -preset veryfast -maxrate 1000k -bufsize 2000k -f flv rtmp://live.twitch.tv/app/streamkey
     ```

## 14. **视频/音频加速/减速**
   - 调整视频或音频的播放速度。
     ```bash
     ffmpeg -i input.mp4 -vf "setpts=0.5*PTS" -af "atempo=2.0" output.mp4
     ```

## 15. **视频/音频拼接**
   - 将多个视频或音频文件拼接为一个文件。
     ```bash
     ffmpeg -f concat -i filelist.txt -c copy output.mp4
     ```

## 16. **视频/音频元数据编辑**
   - 修改视频或音频的元数据信息。
     ```bash
     ffmpeg -i input.mp4 -metadata title="My Video" -metadata author="My Name" output.mp4
     ```

## 17. **视频/音频降噪**
   - 对音频或视频进行降噪处理。
     ```bash
     ffmpeg -i input.mp4 -af "afftdn=nf=-20" output.mp4
     ```

## 18. **视频/音频格式探测**
   - 查看视频或音频文件的详细信息。
     ```bash
     ffmpeg -i input.mp4
     ```

## 19. **视频/音频截图**
   - 从视频中生成缩略图或截图。
     ```bash
     ffmpeg -i input.mp4 -ss 00:00:10 -vframes 1 output.jpg
     ```

## 20. **视频/音频倒放**
   - 将视频或音频倒放。
     ```bash
     ffmpeg -i input.mp4 -vf reverse -af areverse output.mp4
     ```

---

FFmpeg 功能强大且灵活，适用于各种多媒体处理场景。通过命令行参数，可以实现几乎所有的音视频处理需求。